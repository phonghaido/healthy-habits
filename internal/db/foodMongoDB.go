package db

import (
	"context"
	"strings"

	internal_type "github.com/phonghaido/healthy-habits/internal/types"
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

func (f FoodCollection) InsertOne(item internal_type.FoundationFood) error {
	_, err := f.Collection.InsertOne(f.Context, item)
	if err != nil {
		return err
	}
	return nil
}

func (f FoodCollection) InsertMany(items []internal_type.FoundationFood) error {
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

func (f FoodCollection) FindOne(filter interface{}) (internal_type.FoundationFood, error) {
	var result internal_type.FoundationFood
	err := f.Collection.FindOne(f.Context, filter).Decode(&result)
	if err != nil {
		return internal_type.FoundationFood{}, err
	}

	return result, nil
}

func (f FoodCollection) FindMany(reqBody internal_type.FindFoodReqBody) ([]internal_type.FoundationFood, error) {
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
		return nil, nil
	}

	cursor, err := f.Collection.Find(f.Context, filter)
	if err != nil {
		return nil, err
	}

	var foods []internal_type.FoundationFood
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

func (f FoodCollection) UpdateOne(filter, update bson.D) (internal_type.FoundationFood, error) {
	return internal_type.FoundationFood{}, nil
}

func (f FoodCollection) UpdateMany(filter, update bson.D) ([]internal_type.FoundationFood, error) {
	return nil, nil
}

func (f FoodCollection) ReplaceOne(filter bson.D, replacement internal_type.FoundationFood) (internal_type.FoundationFood, error) {
	return internal_type.FoundationFood{}, nil
}
