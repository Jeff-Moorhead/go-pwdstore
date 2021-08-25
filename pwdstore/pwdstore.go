package pwdstore

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/jeff-moorhead/go-pwdmgr/encryption"
)

type PasswordStore struct {
	key       []byte
	passwords map[string]string
}

func NewPasswordStore(passwords, key []byte) (*PasswordStore, error) {
	var pwds map[string]string
	err := json.Unmarshal(passwords, &pwds)
	if err != nil {
		return nil, err
	}

	store := PasswordStore{
		key,
		pwds,
	}
	return &store, nil
}

func (self *PasswordStore) DecodeKey() ([]byte, error) {
	decodedKey, err := encryption.DecodeFromHex(self.key)
	if err != nil {
		return nil, err
	}

	return decodedKey, nil
}

func (self *PasswordStore) Get(key string) (string, error) {
	encrypted, ok := self.passwords[key]
	if !ok {
		return "", fmt.Errorf("Key %q not found", key)
	}

	rawEncrypted, err := encryption.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	decodedKey, err := self.DecodeKey()
	if err != nil {
		return "", err
	}

	decrypted, err := encryption.Decrypt(rawEncrypted, decodedKey)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}

func (self *PasswordStore) Set(key, value string) error {
	decodedKey, err := self.DecodeKey()
	if err != nil {
		return fmt.Errorf("Error occurred decoding key: %q", err)
	}

	bytesValue := []byte(value)

	encrypted, err := encryption.Encrypt(bytesValue, decodedKey)
	if err != nil {
		return fmt.Errorf("Error occurred encrypting value: %q", err)
	}

	encodedPwd := encryption.EncodeToHex(encrypted)
	self.passwords[key] = string(encodedPwd)

	return nil
}

func (self *PasswordStore) Save(writer io.Writer) error {
	b, err := json.Marshal(self.passwords)
	if err != nil {
		return err
	}

	_, err = writer.Write(b)
	if err != nil {
		return err
	}

	return nil
}
