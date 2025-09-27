package hsm

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/pem"
	"errors"
)

// Adapter abstracts HSM operations; uses SoftHSM for local tests.
type Adapter struct{}

func NewAdapter() *Adapter { return &Adapter{} }

func (a *Adapter) GenerateKey(label string, bits int) (*rsa.PrivateKey, error) {
	// Replace with PKCS#11 logic for real HSM.
	return rsa.GenerateKey(rand.Reader, bits)
}

func (a *Adapter) Sign(label string, csrDER []byte) ([]byte, error) {
	// In a real HSM, locate key by label and sign.
	block, _ := pem.Decode(csrDER)
	if block == nil {
		return nil, errors.New("invalid CSR PEM")
	}
	// Mock: echo back CSR
	return csrDER, nil
}
