package pwdstore

import "io"

type Store interface {
	Get(key string) (string, error)
	Set(key, value string) error
	Remove(key string)
	Save(writer io.Writer) error
	Keys() []string
}
