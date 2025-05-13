package repository

import (
	"context"

	"github.com/emaldie/secret-api/internal/server/dto"
	apperrors "github.com/emaldie/secret-api/pkg/errors"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SecretRepository interface {
	Create(ctx context.Context, input dto.CreateSecretDto) (interface{}, error)
	Get(ctx context.Context) (interface{}, error)
}

type secretRepository struct {
	db *mongo.Client
}

func NewSecretRepository(db *mongo.Client) SecretRepository {
	return &secretRepository{db: db}
}

func (s *secretRepository) Create(ctx context.Context, input dto.CreateSecretDto) (interface{}, error) {
	secretsCollection := s.db.Database("secrets-api").Collection("secrets")
	createdSecret, err := secretsCollection.InsertOne(ctx, input)
	if err != nil {
		return "", apperrors.InternalError("Internal server error", err)
	}
	return createdSecret.InsertedID, nil
}

func (s *secretRepository) Get(ctx context.Context) (interface{}, error) {
	return struct{}{}, nil
}
