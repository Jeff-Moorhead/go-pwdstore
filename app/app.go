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
	var err error
	if self.opts.Init {
		fmt.Println("Initializing new password manager...")
		err = self.initializeBackend()
		if err != nil {
			return fmt.Errorf("Something went wrong initializing manager, %v", err)
		}

		fmt.Println("Initialization complete! Run `go-pwdmgr --add --password <password> --title <title>` to add a new password.")
		return nil
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

	err = os.WriteFile(mgrfile, nil, 0600)
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
