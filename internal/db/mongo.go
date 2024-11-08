package db

import (
	"context"

	"github.com/phonghaido/healthy-habits/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Context context.Context
	Client  *mongo.Client
}

func NewMongoClient() (*MongoClient, error) {
	mongodbConfig, err := config.GetMongoDBConfig()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongodbConfig.MongoDBConnStr).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	return &MongoClient{
		Context: ctx,
		Client:  client,
	}, nil
}

func (m *MongoClient) Disconnect() error {
	if err := m.Client.Disconnect(m.Context); err != nil {
		return err
	}
	return nil
}
