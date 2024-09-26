package controller

import (
	"net/http"
)

func RegisterHealthHandlers(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", health)
}

func health(writer http.ResponseWriter, _ *http.Request) {
	_, _ = writer.Write([]byte("OK"))
}
