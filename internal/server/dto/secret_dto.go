package dto

type CreateSecretInput struct {
	Message        string `json:"message" validate:"required"`
	ViewCount      int    `json:"view_count"`
	ExpirationTime int    `json:"exp_time"`
}

type CreateSecretDto struct {
	HashID         string `json:"hash_id"`
	Message        string `json:"message" validate:"required, min=1"`
	ViewCount      int    `json:"view_count"`
	ExpirationTime int    `json:"exp_time"`
}

type SecretResponse CreateSecretInput
