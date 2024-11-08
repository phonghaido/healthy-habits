package db

import (
	"context"
	"strings"

	"github.com/phonghaido/healthy-habits/internal/diet"
	internal_type "github.com/phonghaido/healthy-habits/internal/types"
	custom_error "github.com/phonghaido/healthy-habits/pkg/error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MealCollection struct {
	Context    context.Context
	Collection *mongo.Collection
}

func NewMealCollection(c *MongoClient) *MealCollection {
	coll := c.Client.Database("healthy_habits").Collection("meal")
	return &MealCollection{
		Context:    c.Context,
		Collection: coll,
	}
}

func (m MealCollection) FindMany(reqBody internal_type.FindMealReqBody) ([]diet.MealPlan, error) {
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

	cursor, err := m.Collection.Find(m.Context, filter)
	if err != nil {
		return nil, err
	}

	var meals []diet.MealPlan
	if err = cursor.All(m.Context, &meals); err != nil {
		return nil, err
	}

	return meals, nil
}

func (m MealCollection) InsertOne(plan diet.MealPlan) error {
	plan.TotalNutrients = plan.CalculateTotalNutrients()

	_, err := m.Collection.InsertOne(m.Context, plan)
	if err != nil {
		return err
	}
	return nil
}

func (m MealCollection) UpdateOne(plan diet.MealPlan) error {
	filter := bson.D{{Key: "id", Value: plan.ID}}

	plan.TotalNutrients = plan.CalculateTotalNutrients()

	update := bson.D{
		{Key: "$set", Value: bson.D{{Key: "name", Value: plan.Name}}},
		{Key: "$set", Value: bson.D{{Key: "items", Value: plan.Items}}},
		{Key: "$set", Value: bson.D{{Key: "type", Value: plan.Type}}},
		{Key: "$set", Value: bson.D{{Key: "description", Value: plan.Description}}},
		{Key: "$set", Value: bson.D{{Key: "totalNutrients", Value: plan.TotalNutrients}}},
	}

	_, err := m.Collection.UpdateOne(m.Context, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (m MealCollection) DeleteOne(id string) error {
	filter := bson.D{{Key: "id", Value: id}}

	_, err := m.Collection.DeleteOne(m.Context, filter)
	if err != nil {
		return err
	}
	return nil
}

func (m MealCollection) DeleteMany(ids []string) error {
	var condition bson.A
	for _, v := range ids {
		condition = append(condition, bson.D{{Key: "id", Value: v}})
	}
	filter := bson.D{{Key: "$or", Value: condition}}
	_, err := m.Collection.DeleteMany(m.Context, filter)
	if err != nil {
		return err
	}
	return nil
}
