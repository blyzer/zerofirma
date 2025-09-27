package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/storage/upload", func(w http.ResponseWriter, r *http.Request) {
		data, _ := io.ReadAll(r.Body)
		ct, _ := store.EncryptAndStore(data, []byte("examplekey123456"))
		w.Write(ct)
	})
	http.ListenAndServe(":9090", nil)
}
