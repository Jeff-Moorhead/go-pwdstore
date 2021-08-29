package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jeff-moorhead/go-pwdstore/encryption"
	"github.com/jeff-moorhead/go-pwdstore/pwdstore"
)

func BaseDir() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Could not get default directory: %v", err)
	}

	return filepath.Join(homedir, ".pwdmgr"), nil
}

func DefaultFile(base string) (string, error) {
	dir, err := BaseDir()
	if err != nil {
		return "", fmt.Errorf("Could not get default file: %v", err)
	}

	return filepath.Join(dir, base), nil
}

type App struct {
	manager pwdstore.Store // Program to the interface
	opts    Options
}

func NewApp() (*App, error) {
	options := GetOptions()
	app := App{
		manager: nil,
		opts:    options,
	}
	return &app, nil
}

func (self *App) Run() error {
	if self.opts.Init {
		fmt.Println("Initializing new password manager...")
		err := self.initializeBackend()
		if err != nil {
			return fmt.Errorf("Something went wrong initializing manager, %v", err)
		}

		fmt.Println("Initialization complete! Run `go-pwdstore --add --password <password> --title <title>` to add a new password.")

		// Stop execution after initialization
		return nil
	}

	err := self.loadManager()
	if err != nil {
		return err
	}

	if self.opts.All {
		titles := self.getTitles()

		fmt.Println() // Padding on top
		fmt.Println("Titles")
		fmt.Println("--------------------------------------------------------------------")

		for _, t := range titles {
			fmt.Printf(" - %v\n", t)
		}

		fmt.Println() // Padding on bottom

		// Stop execution after printing titles
		return nil
	}

	switch {

	// Order of priority for app flags: Add, Set, Remove. Only one will run.
	case self.opts.Add:
		ok := self.checkAddArgs()
		if !ok {
			return fmt.Errorf("Missing arguments: to add a password, include --title and --password arguments.")
		}

		err := self.addPassword()
		if err != nil {
			return fmt.Errorf("Something went wrong adding password, %v", err)
		}

		fmt.Println("Password saved!")

	case self.opts.Set:
		ok := self.checkSetArgs()
		if !ok {
			return fmt.Errorf("Missing arguments: to set a password, include --title and --password arguments.")
		}

		err := self.setPassword()
		if err != nil {
			return fmt.Errorf("Something went wrong setting password, %v", err)
		}

		fmt.Println("Password saved!")

	case self.opts.Remove:
		ok := self.checkRemoveArgs()
		if !ok {
			return fmt.Errorf("Missing arguments: to remove a password, include --title argument.")
		}

		err := self.removePassword()
		if err != nil {
			return fmt.Errorf("Something went wrong removing password, %v", err)
		}

		fmt.Println("Password removed!")

	default:
		// Default is to fetch the password
		ok := self.checkGetArgs()
		if !ok {
			return fmt.Errorf("Missing argument: to get a password, include --title argument")
		}

		pwd, err := self.getPassword()
		if err != nil {
			return fmt.Errorf("Something went wrong getting password, %v", err)
		}

		fmt.Printf("Password for %v: %v\n", self.opts.Title, pwd)
	}

	return nil
}

func (self *App) initializeBackend() error {
	err := self.initBaseDir()
	if err != nil {
		return err
	}

	key, err := encryption.NewEncodedKey()
	if err != nil {
		return err
	}

	keyfile, err := self.getKeyFilename()
	if err != nil {
		return err
	}

	err = os.WriteFile(keyfile, key, 0600)
	if err != nil {
		return err
	}

	mgrfile, err := self.getFilename()
	if err != nil {
		return err
	}

	err = os.WriteFile(mgrfile, []byte("{}"), 0600)
	if err != nil {
		return err
	}

	return nil
}

func (self *App) initBaseDir() error {
	// Only create the base directory if the user does not specify a data file or key file
	if self.opts.File == "" || self.opts.Key == "" {
		mgrRoot, err := BaseDir()
		if err != nil {
			return err
		}

		err = os.MkdirAll(mgrRoot, 0700)
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *App) getTitles() []string {
	return self.manager.Keys()
}

func (self *App) getKeyFilename() (string, error) {
	if self.opts.Key == "" {
		return DefaultFile("store.key")
	}
	return self.opts.Key, nil
}

func (self *App) getFilename() (string, error) {
	if self.opts.File == "" {
		return DefaultFile("store.json")
	}
	return self.opts.File, nil
}

func (self *App) loadManager() error {
	keyname, err := self.getKeyFilename()
	if err != nil {
		return err
	}

	key, err := os.ReadFile(keyname)
	if err != nil {
		return err
	}

	filename, err := self.getFilename()
	if err != nil {
		return err
	}

	pwds, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	mgr, err := pwdstore.NewPasswordStore(pwds, key)
	if err != nil {
		return err
	}

	self.manager = mgr
	return nil
}

func (self *App) checkAddArgs() bool {
	// To add add password, password and title are required
	return self.opts.Password != "" && self.opts.Title != ""
}

func (self *App) addPassword() error {
	existing, _ := self.manager.Get(self.opts.Title)
	if existing != "" {
		return fmt.Errorf("Password with title %v already exists. Use --set to change or --remove to remove.", self.opts.Title)
	}

	err := self.manager.Set(self.opts.Title, self.opts.Password)
	if err != nil {
		return nil
	}

	err = self.save()
	if err != nil {
		return err
	}

	return nil
}

func (self *App) checkGetArgs() bool {
	return self.opts.Title != ""
}

func (self *App) getPassword() (string, error) {
	return self.manager.Get(self.opts.Title)
}

func (self *App) checkSetArgs() bool {
	return self.opts.Title != "" && self.opts.Password != ""
}

func (self *App) setPassword() error {
	_, err := self.manager.Get(self.opts.Title)
	if err != nil {
		return err
	}

	err = self.manager.Set(self.opts.Title, self.opts.Password)
	if err != nil {
		return err
	}

	err = self.save()
	if err != nil {
		return err
	}

	return nil
}

func (self *App) checkRemoveArgs() bool {
	return self.opts.Title != ""
}

func (self *App) removePassword() error {
	self.manager.Remove(self.opts.Title)

	err := self.save()
	if err != nil {
		return err
	}

	return nil
}

func (self *App) save() error {
	filename, err := self.getFilename()
	if err != nil {
		return err
	}

	// Open backend file for writing
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer f.Close()
	err = self.manager.Save(f)
	if err != nil {
		return err
	}

	return nil
}
