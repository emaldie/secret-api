package handlers

import (
	"log/slog"

	"net/http"

	"github.com/emaldie/secret-api/internal/server/services"
	"github.com/go-playground/validator"
)

type SecretHandler struct {
	secretService services.SecretService
	logger        *slog.Logger
	validator     *validator.Validate
}

func NewSecretsHandler(
	secretService services.SecretService,
	logger *slog.Logger,
	validator *validator.Validate,
) SecretHandler {
	return SecretHandler{
		secretService: secretService,
		logger:        logger,
		validator:     validator,
	}
}

func (s *SecretHandler) GetSecret(w http.ResponseWriter, r *http.Request) {

}

func (s *SecretHandler) CreateSecret(w http.ResponseWriter, r *http.Request) {

}
