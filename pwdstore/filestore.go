package pwdstore

import "io"

type PasswordStore map[string]string

func NewPasswordStore(reader io.Reader) PasswordStore {
	return PasswordStore{}
}

func (self *PasswordStore) Get(key string) (string, error) {
	return "", nil
}

func (self *PasswordStore) Set(key string) error {
	return nil
}
