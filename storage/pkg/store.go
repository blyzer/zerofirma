package storage

import (
	"crypto/aes"
	"crypto/cipher"
)

func EncryptAndStore(data []byte, key []byte) ([]byte, error) {
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	// write ciphertext to disk or S3
	return ciphertext, nil
}
