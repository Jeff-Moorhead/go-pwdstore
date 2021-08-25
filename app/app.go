package app

import (
	"fmt"
	"os"

	"github.com/jeff-moorhead/go-pwdmgr/encryption"
	"github.com/jeff-moorhead/go-pwdmgr/pwdstore"
)

type App struct {
	manager *pwdstore.PasswordStore
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
	switch {
	case self.opts.Init:
		fmt.Println("Initializing new password manager...")
		err := self.initializeBackend()
		if err != nil {
			return fmt.Errorf("Something went wrong initializing manager, %v", err)
		}

		fmt.Println("Initialization complete! Run `go-pwdmgr --add --password <password> --title <title>` to add a new password.")

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

	// TODO: Implement set, remove, and show-titles functionality

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

func (self *App) checkAddArgs() bool {
	// To add add password, password and title are required
	return self.opts.Password != "" && self.opts.Title != ""
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

func (self *App) addPassword() error {
	err := self.loadManager()
	if err != nil {
		return err
	}

	existing, _ := self.manager.Get(self.opts.Title)
	if existing != "" {
		return fmt.Errorf("Password with title %v already exists. Use --set to change or --remove to remove.", self.opts.Title)
	}

	err = self.manager.Set(self.opts.Title, self.opts.Password)
	if err != nil {
		return nil
	}

	filename, err := self.getFilename()
	if err != nil {
		return err
	}

	// Open back end file for writing
	f, err := os.OpenFile(filename, os.O_RDWR, os.ModeAppend)
	if err != nil {
		return err
	}

	defer f.Close()
	self.manager.Save(f)

	return nil
}

func (self *App) checkGetArgs() bool {
	return self.opts.Title != ""
}

func (self *App) getPassword() (string, error) {
	err := self.loadManager()
	if err != nil {
		return "", err
	}

	return self.manager.Get(self.opts.Title)
}
