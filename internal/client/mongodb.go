package client

import (
	"context"
	"we-connect-test/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDBClient(cfg *config.Cfg) (*mongo.Client, error) {
	uri := cfg.GetString("mongodb.dsn")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.Background(), opts)
	return client, err
}
