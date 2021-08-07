package pwdstore

import (
	"strings"
	"testing"
)

func TestNewPasswordStore(t *testing.T) {
	data := `{"example.com":"foobar","wizbang.org":"i-am-a-teapot","gopher.net":"fizbaz"}`
	reader := strings.NewReader(data)

	expected := PasswordStore{
		"example.com": "foobar",
		"wizbang.org": "i-am-a-teapot",
		"gopher.net":  "fizbaz",
	}

	got := NewPasswordStore(reader)
	for k, v := range expected {
		got_val, ok := got[k]
		if !ok {
			t.Errorf("Key %q is missing from expected", k)
		}
		if got_val != v {
			t.Errorf("Incorrect value in return value: got %q, expected %q", v, expected_val)
		}
	}
}
