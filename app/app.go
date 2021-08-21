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
	if self.opts.Init {
		fmt.Println("Hello!")
		mgrRoot, err := BaseDir()
		if err != nil {
			return fmt.Errorf("Something went wrong initializing: %v", err)
		}

		err = os.MkdirAll(mgrRoot, 0700)
		if err != nil {
			return fmt.Errorf("Something went wrong initializing: %v", err)
		}

		key, err := encryption.NewEncodedKey()
		if err != nil {
			return fmt.Errorf("Something went wrong initializing: %v", err)
		}

		if self.opts.Key == "" {
			keyfile, err := DefaultFile("store.key")
			if err != nil {
				return fmt.Errorf("Something went wrong initializing: %v", err)
			}
			err = os.WriteFile(keyfile, key, 0600)
			if err != nil {
				return fmt.Errorf("Something went wrong initializing: %v", err)
			}
		} else {
			err := os.WriteFile(self.opts.Key, nil, 0600)
			if err != nil {
				return fmt.Errorf("Something went wrong initializing: %v", err)
			}
		}

		if self.opts.File == "" {
			mgrfile, err := DefaultFile("store.json")
			if err != nil {
				return fmt.Errorf("Something went wrong initializing: %v", err)
			}
			err = os.WriteFile(mgrfile, nil, 0600)
			if err != nil {
				return fmt.Errorf("Something went wrong initializing: %v", err)
			}
		} else {
			err := os.WriteFile(self.opts.File, nil, 0600)
			if err != nil {
				return fmt.Errorf("Something went wrong initializing: %v", err)
			}
		}
	}

	return nil
}
