// Credit
package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
)

const EncryptionKeySize = 32 // 32 bytes for AES-256

func EncodeToHex(data []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(data)))
	_ = hex.Encode(dst, data)
	return dst
}

func DecodeFromHex(data []byte) ([]byte, error) {
	dst := make([]byte, hex.DecodedLen(len(data)))
	_, err := hex.Decode(dst, data)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func DecodeString(data string) ([]byte, error) {
	return hex.DecodeString(data)
}

// NewEncryptionKey generates a random 256-bit key for Encrypt() and Decrypt().
func NewEncryptionKey() ([]byte, error) {
	key := make([]byte, EncryptionKeySize)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func Encrypt(plaintext []byte, key []byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func Decrypt(ciphertext []byte, key []byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return nil, errors.New("Malformed ciphertext")
	}

	nonce := ciphertext[:gcm.NonceSize()]
	data := ciphertext[gcm.NonceSize():]
	return gcm.Open(nil,
		nonce,
		data,
		nil,
	)
}
