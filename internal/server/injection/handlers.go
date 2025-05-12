package injection

import (
	"log/slog"

	"os"

	"github.com/emaldie/secret-api/internal/server/handlers"
	"github.com/go-playground/validator"
)

type Handlers struct {
	SecretHandler handlers.SecretHandler
}

func InitHandlers(
	services *Services,
	logger *slog.Logger,
	validate *validator.Validate,
) *Handlers {
	if services == nil {
		slog.Error("Error initializing repositories. Repositories mustn't be nil")
		os.Exit(1)
	}
	if logger == nil {
		slog.Error("Error initializing repositories. Logger mustn't be nil")
		os.Exit(1)
	}
	if validate == nil {
		slog.Error("Error initializing repositories. Validator mustn't be nil")
		os.Exit(1)
	}
	secretHandler := handlers.NewSecretsHandler(services.SecretService, logger, validate)

	return &Handlers{
		SecretHandler: secretHandler,
	}
}
