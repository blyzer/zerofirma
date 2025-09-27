package issuance

import (
	"encoding/json"
	"net/http"
)

type CSRRequest struct {
	CSR string `json:"csr_pem"`
}

func (s *Service) HandleCSR(w http.ResponseWriter, r *http.Request) {
	var req CSRRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	cert, err := s.IssueCertificate(r.Context(), []byte(req.CSR))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(cert)
}
