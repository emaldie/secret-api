package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SecretRepository interface {
	Create(ctx context.Context) error
	Get(ctx context.Context) error
}

type secretRepository struct {
	db *mongo.Client
}

func NewSecretRepository(db *mongo.Client) SecretRepository {
	return &secretRepository{db: db}
}

func (s *secretRepository) Create(ctx context.Context) error {
	return nil
}

func (s *secretRepository) Get(ctx context.Context) error {
	return nil
}
