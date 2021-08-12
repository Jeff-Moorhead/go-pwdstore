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
		"foobar",
		map[string]string{
			"example.com": "foobar",
			"wizbang.org": "i-am-a-teapot",
			"gopher.net":  "fizbaz",
		},
	}

	got, err := NewPasswordStore(reader)
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

// TODO: test Get and Set methods on passwordstore object
