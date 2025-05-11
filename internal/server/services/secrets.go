package services

import (
	"context"

	"github.com/emaldie/secret-api/internal/server/repository"
	"github.com/go-playground/validator"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SecretService interface {
	CreateTemporary(ctx context.Context)
	CreateOnetime(ctx context.Context)
	GetTemp(ctx context.Context)
	GetOnetime(ctx context.Context)
}

type secretService struct {
	secretRepo repository.SecretRepository
	validator  *validator.Validate
	db         *mongo.Client
	rdb        *redis.Client
}

func NewSecretService(
	secretRepo repository.SecretRepository,
	validator *validator.Validate,
	db *mongo.Client,
	rdb *redis.Client) *secretService {
	return &secretService{
		secretRepo: secretRepo,
		validator:  validator,
		db:         db,
		rdb:        rdb,
	}
}
