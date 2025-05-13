package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SecretMongoModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	HashId    string             `bson:"hash_id,omitempty"`
	Message   string             `bson:"message,omitempty"`
	ViewCount int                `bson:"view_count,omitempty"`
}

type SecretRedisModel struct {
	Message        string
	ViewCount      int
	ExpirationTime time.Duration
	InsertionTime  time.Time
}

type SecretModelInterface interface {
}
