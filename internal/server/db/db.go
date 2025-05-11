package db

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/emaldie/secret-api/internal/server/config"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func InitMongo(cfg *config.MongoConfig) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(cfg.Uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, fmt.Errorf("error connecting to db: %w", err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			slog.Error("error pinging db", "error", err)
		}
	}()

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, fmt.Errorf("error pinging db: %w", err)
	}
	return client, nil
}

func InitRedis(cfg *config.RedisConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Username: cfg.Username,
		Password: cfg.Password,
		DB:       cfg.Db,
	})

	return rdb, nil
}
