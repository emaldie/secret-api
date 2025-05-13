package services

import (
	"context"
	"encoding/base64"

	"github.com/emaldie/secret-api/internal/server/dto"
	"github.com/emaldie/secret-api/internal/server/models"
	"github.com/emaldie/secret-api/internal/server/repository"
	"github.com/go-playground/validator"
	"github.com/redis/go-redis/v9"

	"crypto/rand"
)

type SecretService interface {
	CreateSecret(ctx context.Context, input dto.CreateSecretInput) (interface{}, error)
	GetSecret(ctx context.Context, hashId string) (models.SecretModelInterface, error)
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

func (s *secretService) CreateSecret(ctx context.Context, input dto.CreateSecretInput) (interface{}, error) {

	code := make([]byte, 10)
	_, err := rand.Read(code)
	if err != nil {
		return struct{}{}, err
	}
	var secret = dto.CreateSecretDto{
		HashID:         base64.URLEncoding.EncodeToString(code),
		Message:        input.Message,
		ExpirationTime: input.ExpirationTime,
		ViewCount:      input.ViewCount,
	}

	response, err := s.secretRepo.Create(ctx, secret)
	if err != nil {
		return struct{}{}, err
	}
	return response, nil
}

func (s *secretService) GetSecret(ctx context.Context, hashId string) (models.SecretModelInterface, error) {
	return struct{}{}, nil
}
