package pwdstore

type Store interface {
	Get(key string) (string, error)
	Set(key, value string) error
}
