package injection

import (
	"log/slog"
	"os"

	"github.com/emaldie/secret-api/internal/server/config"
	"github.com/emaldie/secret-api/internal/server/services"
	"github.com/go-playground/validator"
	"github.com/redis/go-redis/v9"
)

type Services struct {
	SecretService services.SecretService
}

func InitServices(
	repos *Repositories,
	validate *validator.Validate,
	rdb *redis.Client,
	config *config.AppConfig,
) *Services {
	if repos == nil {
		slog.Error("Error initializing repositories. Repositories mustn't be nil")
		os.Exit(1)
	}
	if validate == nil {
		slog.Error("Error initializing repositories. Validator mustn't be nil")
		os.Exit(1)
	}
	if rdb == nil {
		slog.Error("Error initializing repositories. Redis client mustn't be nil")
		os.Exit(1)
	}
	if config == nil {
		slog.Error("Error initializing repositories. Config mustn't be nil")
		os.Exit(1)
	}

	secretService := services.NewSecretService(repos.SecretRepository, validate, rdb)

	return &Services{SecretService: secretService}
}
