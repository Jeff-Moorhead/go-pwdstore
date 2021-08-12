package pwdstore

import (
	"encoding/json"
	"io"
)

type PasswordStore struct {
	key       string
	passwords map[string]string
}

func NewPasswordStore(reader io.Reader) (*PasswordStore, error) {
	var pwds map[string]string
	err := json.NewDecoder(reader).Decode(&pwds)
	if err != nil {
		return nil, err
	}

	store := PasswordStore{
		"foobar", // TODO: Generalize this field
		pwds,
	}
	return &store, nil
}

func (self *PasswordStore) Get(key string) (string, error) {
	// TODO: Implement
	return "", nil
}

func (self *PasswordStore) Set(key string) error {
	// TODO: Implement
	return nil
}
