package injection

import (
	"log/slog"
	"os"

	"github.com/emaldie/secret-api/internal/server/repository"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repositories struct {
	SecretRepository repository.SecretRepository
}

func InitRepositories(db *mongo.Client) *Repositories {
	if db == nil {
		slog.Error("Error initializing repositories. Database mustn't be nil")
		os.Exit(1)
	}

	secretRepository := repository.NewSecretRepository(db)
	return &Repositories{
		SecretRepository: secretRepository,
	}
}
