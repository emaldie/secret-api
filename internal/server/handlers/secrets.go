package handlers

import (
	"log/slog"

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
) *SecretHandler {
	return &SecretHandler{secretService: secretService, logger: logger, validator: validator}
}
