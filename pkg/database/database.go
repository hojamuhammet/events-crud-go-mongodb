package database

import (
	"context"
	"events/internal/config"
	"events/pkg/lib/utils"
	"log"
	"log/slog"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DBClient   *mongo.Client
	DBDatabase *mongo.Database
)

func InitDB(cfg *config.Config) error {
	clientOptions := options.Client().ApplyURI(cfg.MongoDB.URI)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		slog.Error("Error connecting to MongoDB: %v", utils.Err(err))
		return err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error pinging MongoDB: %v", utils.Err(err))
		return err
	}

	DBClient = client

	DBDatabase = client.Database(cfg.MongoDB.Database)

	return nil
}

func Close() {
	if DBClient != nil {
		if err := DBClient.Disconnect(context.Background()); err != nil {
			slog.Info("Error closing MongoDB connection: %v", utils.Err(err))
		} else {
			slog.Info("MongoDB connection closed")
		}
	}
}

func GetDB() *mongo.Database {
	return DBDatabase
}
