package db

import (
	"context"
	"strings"

	internal_type "github.com/phonghaido/healthy-habits/internal"
	"github.com/phonghaido/healthy-habits/internal/config"
	"github.com/phonghaido/healthy-habits/internal/diet"
	custom_error "github.com/phonghaido/healthy-habits/pkg/error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MealMongoDBClient *MongoDBMealClient

type MongoDBMealClient struct {
	Context    context.Context
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewMongoDBFMealClient() (*MongoDBMealClient, error) {
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

	coll := client.Database("healthy_habits").Collection("meal")

	return &MongoDBMealClient{
		Context:    ctx,
		Client:     client,
		Collection: coll,
	}, nil
}

func (c MongoDBMealClient) Disconnect() error {
	if err := c.Client.Disconnect(c.Context); err != nil {
		return err
	}
	return nil
}

func (c MongoDBMealClient) FindMany(reqBody internal_type.FindMealReqBody) ([]diet.MealPlan, error) {
	var filter bson.D
	if reqBody.Name == "" {
		return nil, custom_error.InvalidRequestBody("name")
	}

	words := strings.Fields(reqBody.Name)
	conditions := bson.A{}

	for _, word := range words {
		conditions = append(conditions, bson.D{
			{Key: "name", Value: bson.D{
				{Key: "$regex", Value: word},
				{Key: "$options", Value: "i"},
			}},
		})
	}

	filter = bson.D{{Key: "$and", Value: conditions}}

	cursor, err := c.Collection.Find(c.Context, filter)
	if err != nil {
		return nil, err
	}

	var meals []diet.MealPlan
	if err = cursor.All(c.Context, &meals); err != nil {
		return nil, err
	}

	return meals, nil
}

func (c MongoDBMealClient) InsertOne(plan diet.MealPlan) error {
	plan.TotalNutrients = plan.CalculateTotalNutrients()

	_, err := c.Collection.InsertOne(c.Context, plan)
	if err != nil {
		return err
	}
	return nil
}

func (c MongoDBMealClient) UpdateOne(plan diet.MealPlan) error {
	filter := bson.D{{Key: "id", Value: plan.ID}}

	plan.TotalNutrients = plan.CalculateTotalNutrients()

	update := bson.D{
		{Key: "$set", Value: bson.D{{Key: "name", Value: plan.Name}}},
		{Key: "$set", Value: bson.D{{Key: "items", Value: plan.Items}}},
		{Key: "$set", Value: bson.D{{Key: "type", Value: plan.Type}}},
		{Key: "$set", Value: bson.D{{Key: "description", Value: plan.Description}}},
		{Key: "$set", Value: bson.D{{Key: "totalNutrients", Value: plan.TotalNutrients}}},
	}

	_, err := c.Collection.UpdateOne(c.Context, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (c MongoDBMealClient) DeleteOne(id string) error {
	filter := bson.D{{Key: "id", Value: id}}

	_, err := c.Collection.DeleteOne(c.Context, filter)
	if err != nil {
		return err
	}
	return nil
}

func (c MongoDBMealClient) DeleteMany(ids []string) error {
	var condition bson.A
	for _, v := range ids {
		condition = append(condition, bson.D{{Key: "id", Value: v}})
	}
	filter := bson.D{{Key: "$or", Value: condition}}
	_, err := c.Collection.DeleteMany(c.Context, filter)
	if err != nil {
		return err
	}
	return nil
}
