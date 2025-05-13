package handlers

import (
	"log/slog"

	"net/http"

	"github.com/emaldie/secret-api/internal/server/dto"
	"github.com/emaldie/secret-api/internal/server/services"
	"github.com/emaldie/secret-api/pkg/errors"
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
	// hashId := r.PathValue("id")

	// secret, err := s.secretService.GetSecret(r.Context(), hashId)
}

func (s *SecretHandler) CreateSecret(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateSecretInput
	if err := BindJSON(r, &input, func(v any) error {
		return s.validator.Struct(v)
	}); err != nil {
		RespondError(w, err)
		return
	}

	secret, err := s.secretService.CreateSecret(r.Context(), input)
	if err != nil {
		RespondError(w, err)
		return
	}

	RespondJSON(w, http.StatusCreated, secret, errors.Error{})
}
