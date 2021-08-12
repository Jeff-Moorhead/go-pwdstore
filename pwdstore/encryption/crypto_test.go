package encryption

import "testing"

var (
	encryptionKey = make([]byte, 32)
	plaintext     = []byte("Hello, this is a test.")

	encrypted []byte
)

func setUp() {
	encryptionKey, _ = NewEncryptionKey()
}
func TestNewEncryptionKey(t *testing.T) {
	k, err := NewEncryptionKey()
	if err != nil {
		t.Fatalf("An unexpected error occurred in NewEncryptionKey: %q", err)
	}

	keySize := len(k)
	if keySize != EncryptionKeySize {
		t.Errorf("Unexpected key size: expected %v, got %v", EncryptionKeySize, keySize)
	}
}

func TestEncrypt(t *testing.T) {
	var err error
	encrypted, err = Encrypt(plaintext, encryptionKey)
	if err != nil {
		t.Fatalf("An unexpected error occurred in Encrypt: %q", err)
	}
}

func TestDecrypt(t *testing.T) {
	// Global encrypted should decrypt to "Hello, this is a test.".
	got, err := Decrypt(encrypted, encryptionKey)
	if err != nil {
		t.Fatalf("An unexpected error occurred in Decrypt: %q", err)
	}

	if string(got) != string(plaintext) {
		t.Errorf("Ciphertext did not decrypt properly: got %q, expected %q", string(got), string(plaintext))
	}
}

func TestMain(m *testing.M) {
	setUp()
	m.Run()
}
