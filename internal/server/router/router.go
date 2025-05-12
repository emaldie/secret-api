package api

import (
	"net/http"

	"github.com/emaldie/secret-api/internal/server/handlers"
)

type RouterConfig struct {
	SecretHandler *handlers.SecretHandler
}

func Setup(mux *http.ServeMux, config RouterConfig) {
	mux.HandleFunc("GET /s/{id}", config.SecretHandler.GetSecret)
	mux.HandleFunc("POST /", config.SecretHandler.CreateSecret)
}
