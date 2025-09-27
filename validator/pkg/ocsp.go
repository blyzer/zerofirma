package validator

import (
	"bytes"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/crypto/ocsp"
)

// OCSPStatus encapsulates the parsed OCSP response.
type OCSPStatus struct {
	Status     string    // "GOOD", "REVOKED", "UNKNOWN", or "ERROR"
	RevokedAt  time.Time // only for REVOKED
	ThisUpdate time.Time
	NextUpdate time.Time
}

// CheckOCSP issues an OCSP request for certDER against issuerDER at endpoint.
// Returns parsed OCSPStatus or an error if the operation could not complete.
func CheckOCSP(certDER, issuerDER []byte, endpoint string) (OCSPStatus, error) {
	// Parse certificates
	cert, err := x509.ParseCertificate(certDER)
	if err != nil {
		return OCSPStatus{Status: "ERROR"}, fmt.Errorf("parse cert: %w", err)
	}
	issuer, err := x509.ParseCertificate(issuerDER)
	if err != nil {
		return OCSPStatus{Status: "ERROR"}, fmt.Errorf("parse issuer: %w", err)
	}

	// Create OCSP request
	reqBytes, err := ocsp.CreateRequest(cert, issuer, &ocsp.RequestOptions{})
	if err != nil {
		return OCSPStatus{Status: "ERROR"}, fmt.Errorf("create OCSP request: %w", err)
	}

	// Send HTTP POST to OCSP responder
	httpReq, err := http.NewRequest("POST", endpoint, bytes.NewReader(reqBytes))
	if err != nil {
		return OCSPStatus{Status: "ERROR"}, fmt.Errorf("build HTTP request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/ocsp-request")
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return OCSPStatus{Status: "ERROR"}, fmt.Errorf("OCSP POST error: %w", err)
	}
	defer resp.Body.Close()

	// Read and parse OCSP response
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return OCSPStatus{Status: "ERROR"}, fmt.Errorf("read OCSP response: %w", err)
	}
	ocspResp, err := ocsp.ParseResponse(respBytes, issuer)
	if err != nil {
		return OCSPStatus{Status: "ERROR"}, fmt.Errorf("parse OCSP response: %w", err)
	}

	// Map status code to string
	status := "UNKNOWN"
	switch ocspResp.Status {
	case ocsp.Good:
		status = "GOOD"
	case ocsp.Revoked:
		status = "REVOKED"
	case ocsp.Unknown:
		status = "UNKNOWN"
	}

	return OCSPStatus{
		Status:     status,
		RevokedAt:  ocspResp.RevokedAt,
		ThisUpdate: ocspResp.ThisUpdate,
		NextUpdate: ocspResp.NextUpdate,
	}, nil
}
