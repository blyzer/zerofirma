package validator

import (
	"crypto/x509"
	"errors"
	"time"
)

func ValidateChain(chain []*x509.Certificate) error {
	// Simplified: check leaf validity.
	leaf := chain[0]
	if leaf.NotAfter.Before(time.Now()) {
		return errors.New("certificate expired")
	}
	return nil
}
