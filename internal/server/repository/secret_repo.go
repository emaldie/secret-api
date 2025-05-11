package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SecretRepository interface {
	Create(ctx context.Context)
	GetById(ctx context.Context)
}

type secretRepository struct {
	db *mongo.Client
}

func NewSecretRepository(db *mongo.Client) *secretRepository {
	return &secretRepository{db: db}
}
