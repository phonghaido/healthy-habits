package db

import (
	"context"
	"fmt"

	"github.com/phonghaido/healthy-habits/internal/config"
	"github.com/phonghaido/healthy-habits/internal/nutrient"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClient struct {
	Context      context.Context
	Collection   string
	DatabaseName string
	MongoClient  *mongo.Client
}

func NewMongoDBClient(coll string) (*MongoDBClient, error) {
	mongodbConfig, err := config.GetMongoDBConfig()
	if err != nil {
		return nil, err
	}

	connStr := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		mongodbConfig.MongoDBUsername,
		mongodbConfig.MongoDBPassword,
		mongodbConfig.MongoDBHost,
		mongodbConfig.MongoDBPort,
	)

	ctx := context.Background()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connStr).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	return &MongoDBClient{
		Context:      ctx,
		Collection:   coll,
		DatabaseName: mongodbConfig.MongoDBDatabase,
		MongoClient:  client,
	}, nil
}

func (c MongoDBClient) DisconnectFromMongoDB() error {
	if err := c.MongoClient.Disconnect(c.Context); err != nil {
		return err
	}
	return nil
}

func (c MongoDBClient) InsertOne(item nutrient.FoodItem) error {
	coll := c.MongoClient.Database(c.DatabaseName).Collection(c.Collection)

	_, err := coll.InsertOne(c.Context, item)
	if err != nil {
		return err
	}
	return nil
}

func (c MongoDBClient) InsertMany(items []any) error {
	coll := c.MongoClient.Database(c.DatabaseName).Collection(c.Collection)

	_, err := coll.InsertMany(c.Context, items)
	if err != nil {
		return err
	}

	return nil
}
