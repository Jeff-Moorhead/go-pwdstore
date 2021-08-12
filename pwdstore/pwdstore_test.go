package pwdstore

import (
	"strings"
	"testing"
)

func TestNewPasswordStore(t *testing.T) {
	passwords := `{"example.com":"foobar","wizbang.org":"i-am-a-teapot","gopher.net":"fizbaz"}`

	// Because NewPasswordStore takes an interface, we are able to easily mock the input using any
	// type that implements io.Reader
	reader := strings.NewReader(passwords)

	expected := PasswordStore{
		[]byte("foobar"),
		map[string]string{
			"example.com": "foobar",
			"wizbang.org": "i-am-a-teapot",
			"gopher.net":  "fizbaz",
		},
	}

	got, err := NewPasswordStore(reader, []byte("foobar"))
	if err != nil {
		t.Fatalf("An error occurred building the password store: %q", err)
	}

	for k, v := range expected.passwords {
		got_val, ok := got.passwords[k]
		if !ok {
			t.Errorf("Key %q is missing from expected", k)
		}
		if got_val != v {
			t.Errorf("Incorrect value in return value: got %q, expected %q", got_val, v)
		}
	}
}

func TestGet(t *testing.T) {
	key := []byte("0775b6eddfdb5aa891490c9ec7530a605787dc45fe930f33e1cf4b5cdb76ed95")
	store := PasswordStore{
		key,
		map[string]string{
			"example.com": "f67430399130705f46f6d605aeb519b8973e50414e3ac7aa986ad1a2d710df8e026d",
			"wizbang.org": "7452b5a9d8dc3940a22ac7149a1f4d80faac04b7429a55785e05b6cbd1628810bab1f345bbf64ecb16",
			"gopher.net":  "8924be134e7f37a488ba6d56e2a78f996fee0e6c60dc5850bb56c432eca136c42037",
		},
	}

	decryptedPwds := map[string]string{"example.com": "foobar", "wizbang.org": "i-am-a-teapot", "gopher.net": "fizbaz"}
	for k := range store.passwords {
		got, err := store.Get(k)
		if err != nil {
			t.Fatalf("An unexpected error occurred in PasswordStore.Get: %q", err)
		}

		expected := decryptedPwds[k]
		if got != expected {
			t.Errorf("Incorrect decrypted value: got %v, expected %v", got, expected)
		}
	}
}

func TestSet(t *testing.T) {
	key := []byte("0775b6eddfdb5aa891490c9ec7530a605787dc45fe930f33e1cf4b5cdb76ed95")
	store := PasswordStore{
		key,
		map[string]string{},
	}
	website := "bigalo.com"
	password := "mypassword"

	err := store.Set(website, password)
	if err != nil {
		t.Fatalf("An unexpected error occurred in PasswordStore.Set: %q", err)
	}

	if store.passwords[website] == password {
		t.Fatalf("Password was not encrypted after it was added.")
	}

	got, _ := store.Get(website)
	if got != password {
		t.Errorf("Decrypting password failed, indicating encryption failed: got %v, expected %v", got, password)
	}
}
