package issuance

import (
	"context"
	"crypto/x509"
	"encoding/pem"
)

type Service struct {
	hsm *hsm.Adapter
}

func NewService(h *hsm.Adapter) *Service {
	return &Service{hsm: h}
}

func (s *Service) IssueCertificate(ctx context.Context, csrPEM []byte) ([]byte, error) {
	block, _ := pem.Decode(csrPEM)
	csr, err := x509.ParseCertificateRequest(block.Bytes)
	if err != nil {
		return nil, err
	}
	return s.hsm.Sign("subca-key", csr.Raw)
}
