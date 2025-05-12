package services

import (
	"context"

	"github.com/emaldie/secret-api/internal/server/repository"
	"github.com/go-playground/validator"
	"github.com/redis/go-redis/v9"
)

type SecretService interface {
	CreateSecret(ctx context.Context) error
	GetSecret(ctx context.Context) error
}

type secretService struct {
	secretRepo repository.SecretRepository
	validator  *validator.Validate
	rdb        *redis.Client
}

func NewSecretService(
	secretRepo repository.SecretRepository,
	validator *validator.Validate,
	rdb *redis.Client) SecretService {
	return &secretService{
		secretRepo: secretRepo,
		validator:  validator,
		rdb:        rdb,
	}
}

func (s *secretService) CreateSecret(ctx context.Context) error {
	return nil
}

func (s *secretService) GetSecret(ctx context.Context) error {
	return nil
}
