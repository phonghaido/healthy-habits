package db

import (
	"context"
	"strings"

	internal_type "github.com/phonghaido/healthy-habits/internal/types"
	"github.com/phonghaido/healthy-habits/internal/usda"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FoodCollection struct {
	Context    context.Context
	Collection *mongo.Collection
}

func NewFoodCollection(c *MongoClient) *FoodCollection {
	coll := c.Client.Database("healthy_habits").Collection("food")
	return &FoodCollection{
		Context:    c.Context,
		Collection: coll,
	}
}

func (f FoodCollection) InsertOne(item usda.FoundationFood) error {
	_, err := f.Collection.InsertOne(f.Context, item)
	if err != nil {
		return err
	}
	return nil
}

func (f FoodCollection) InsertMany(items []usda.FoundationFood) error {
	itemsToAdd := make([]interface{}, len(items))
	for i, v := range items {
		itemsToAdd[i] = v
	}

	_, err := f.Collection.InsertMany(f.Context, itemsToAdd)
	if err != nil {
		return err
	}

	return nil
}

func (f FoodCollection) FindOne(filter interface{}) (usda.FoundationFood, error) {
	var result usda.FoundationFood
	err := f.Collection.FindOne(f.Context, filter).Decode(&result)
	if err != nil {
		return usda.FoundationFood{}, err
	}

	return result, nil
}

func (f FoodCollection) FindMany(reqBody internal_type.FindFoodReqBody) ([]usda.FoundationFood, error) {
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

	cursor, err := f.Collection.Find(f.Context, filter)
	if err != nil {
		return nil, err
	}

	var foods []usda.FoundationFood
	if err = cursor.All(f.Context, &foods); err != nil {
		return nil, err
	}

	return foods, nil
}

func (f FoodCollection) DeleteOne(filter bson.D) error {
	return nil
}

func (f FoodCollection) DeleteMany(filter bson.D) error {
	return nil
}

func (f FoodCollection) UpdateOne(filter, update bson.D) (usda.FoundationFood, error) {
	return usda.FoundationFood{}, nil
}

func (f FoodCollection) UpdateMany(filter, update bson.D) ([]usda.FoundationFood, error) {
	return nil, nil
}

func (f FoodCollection) ReplaceOne(filter bson.D, replacement usda.FoundationFood) (usda.FoundationFood, error) {
	return usda.FoundationFood{}, nil
}
