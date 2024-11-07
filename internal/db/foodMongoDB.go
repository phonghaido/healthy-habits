package db

import (
	"context"
	"strings"

	internal_type "github.com/phonghaido/healthy-habits/internal"
	"github.com/phonghaido/healthy-habits/internal/config"
	"github.com/phonghaido/healthy-habits/internal/usda"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var FoodMongoDBClient *MongoDBFoodClient

type MongoDBFoodClient struct {
	Context    context.Context
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewMongoDBFoodClient() (*MongoDBFoodClient, error) {
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

	coll := client.Database("healthy_habits").Collection("food")

	return &MongoDBFoodClient{
		Context:    ctx,
		Client:     client,
		Collection: coll,
	}, nil
}

func (c MongoDBFoodClient) Disconnect() error {
	if err := c.Client.Disconnect(c.Context); err != nil {
		return err
	}
	return nil
}

func (c MongoDBFoodClient) InsertOne(item usda.FoundationFood) error {
	_, err := c.Collection.InsertOne(c.Context, item)
	if err != nil {
		return err
	}
	return nil
}

func (c MongoDBFoodClient) InsertMany(items []usda.FoundationFood) error {
	itemsToAdd := make([]interface{}, len(items))
	for i, v := range items {
		itemsToAdd[i] = v
	}

	_, err := c.Collection.InsertMany(c.Context, itemsToAdd)
	if err != nil {
		return err
	}

	return nil
}

func (c MongoDBFoodClient) FindOne(filter interface{}) (usda.FoundationFood, error) {
	var result usda.FoundationFood
	err := c.Collection.FindOne(c.Context, filter).Decode(&result)
	if err != nil {
		return usda.FoundationFood{}, err
	}

	return result, nil
}

func (c MongoDBFoodClient) FindMany(reqBody internal_type.FindFoodReqBody) ([]usda.FoundationFood, error) {
	var filter bson.D

	if reqBody.Description != "" {
		words := strings.Fields(reqBody.Description)
		conditions := bson.A{}

		for _, word := range words {
			conditions = append(conditions, bson.D{
				{Key: "description", Value: bson.D{
					{Key: "$regex", Value: word},
					{Key: "$options", Value: "i"},
				}},
			})
		}

		desFilter := bson.D{{Key: "$and", Value: conditions}}

		if reqBody.Category != "" {
			catFilter := bson.D{{Key: "foodCategory.description", Value: reqBody.Category}}
			filter = bson.D{{Key: "$and", Value: bson.A{desFilter, catFilter}}}
		} else {
			filter = desFilter
		}
	} else if reqBody.Category != "" {
		filter = bson.D{{Key: "foodCategory.description", Value: reqBody.Category}}
	} else {
		filter = bson.D{}
	}

	cursor, err := c.Collection.Find(c.Context, filter)
	if err != nil {
		return nil, err
	}

	var foods []usda.FoundationFood
	if err = cursor.All(c.Context, &foods); err != nil {
		return nil, err
	}

	return foods, nil
}

func (c MongoDBFoodClient) DeleteOne(filter bson.D) error {
	return nil
}

func (c MongoDBFoodClient) DeleteMany(filter bson.D) error {
	return nil
}

func (c MongoDBFoodClient) UpdateOne(filter, update bson.D) (usda.FoundationFood, error) {
	return usda.FoundationFood{}, nil
}

func (c MongoDBFoodClient) UpdateMany(filter, update bson.D) ([]usda.FoundationFood, error) {
	return nil, nil
}

func (c MongoDBFoodClient) ReplaceOne(filter bson.D, replacement usda.FoundationFood) (usda.FoundationFood, error) {
	return usda.FoundationFood{}, nil
}
