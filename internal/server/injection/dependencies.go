package injection

import (
	"log/slog"

	"github.com/go-playground/validator"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/emaldie/secret-api/internal/server/config"
)

type Dependencies struct {
	Repositories *Repositories

	Services *Services

	Handlers *Handlers

	Config *config.AppConfig

	Infrastructure Infrastructure
}

type Infrastructure struct {
	DB        *mongo.Client
	Redis     *redis.Client
	Validator *validator.Validate
	Logger    *slog.Logger
}

func NewDependencies(
	db *mongo.Client,
	rdb *redis.Client,
	validate *validator.Validate,
	config *config.AppConfig,
	logger *slog.Logger,
) *Dependencies {
	deps := &Dependencies{
		Config: config,
		Infrastructure: Infrastructure{
			DB:        db,
			Redis:     rdb,
			Validator: validate,
			Logger:    logger,
		},
	}

	deps.Repositories = InitRepositories(db)

	deps.Services = InitServices(deps.Repositories, validate, rdb, config)

	deps.Handlers = InitHandlers(deps.Services, logger, validate)

	return deps
}
