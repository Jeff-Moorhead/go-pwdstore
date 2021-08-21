package app

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type Options struct {
	Init     bool
	File     string
	Key      string
	Name     string
	Password string
	Add      bool
}

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

func GetOptions() Options {
	init := flag.Bool("init", false, "Initialize a new manager and create a new encryption key")
	file := flag.String("file", "", "Path to the password file, default is ~/.pwdmgr/store.json")
	key := flag.String("key", "", "Path to the password key file, default is ~/.pwdmgr/store.key")
	name := flag.String("name", "", "Name of the website, e.g. www.facebook.com")
	password := flag.String("password", "", "Password, only used when adding or modifying")
	add := flag.Bool("add", false, "Set when adding a new password")
	flag.Parse()

	return Options{
		Init:     *init,
		File:     *file,
		Key:      *key,
		Name:     *name,
		Password: *password,
		Add:      *add,
	}
}
