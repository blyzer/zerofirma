package main

import (
	"io"
	"log"
	"net/http"

	"digito-platform/ca-service/internal/hsm"
	"digito-platform/ca-service/internal/issuance"

	"github.com/go-chi/chi/v5"
)

func main() {
	h := hsm.NewAdapter()
	svc := issuance.NewService(h)

	r := chi.NewRouter()
	r.Post("/ca/subca1/csr", func(w http.ResponseWriter, r *http.Request) {
		csrPEM, _ := io.ReadAll(r.Body)
		certDER, err := svc.IssueCertificate(r.Context(), csrPEM)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write(certDER)
	})

	log.Println("CA server listening on :8080")
	http.ListenAndServe(":8080", r)
}
