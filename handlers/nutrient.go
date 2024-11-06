package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/phonghaido/healthy-habits/internal/db"
	errorWrapper "github.com/phonghaido/healthy-habits/pkg/error"
	"go.mongodb.org/mongo-driver/bson"
)

type RequestCategoryBody struct {
	Category string `json:"category,omitempty"`
}

type RequestDescriptionBody struct {
	Description string `json:"description,omitempty"`
}

func HandleGETFindFood(c echo.Context) error {
	result, err := db.FoodMongoDBClient.FindMany(bson.D{{}})
	if err != nil {
		return err
	}
	return errorWrapper.WriteJSON(c, http.StatusOK, result)
}

func HandlePOSTFindFoodByCategory(c echo.Context) error {
	var requestBody RequestCategoryBody
	if err := json.NewDecoder(c.Request().Body).Decode(&requestBody); err != nil {
		return err
	}

	if requestBody.Category == "" {
		return errorWrapper.InvalidRequestBody()
	}

	filter := bson.D{{Key: "foodCategory.description", Value: requestBody.Category}}
	result, err := db.FoodMongoDBClient.FindMany(filter)
	if err != nil {
		return err
	}
	return errorWrapper.WriteJSON(c, http.StatusOK, result)
}

func HandlePOSTFindFoodByDescription(c echo.Context) error {
	var requestBody RequestDescriptionBody
	if err := json.NewDecoder(c.Request().Body).Decode(&requestBody); err != nil {
		return err
	}

	if requestBody.Description == "" {
		return errorWrapper.InvalidRequestBody()
	}

	filter := bson.D{{Key: "description", Value: bson.D{{Key: "$regex", Value: requestBody.Description}}}}
	result, err := db.FoodMongoDBClient.FindMany(filter)
	if err != nil {
		return err
	}

	x := make([]string, len(result))
	for i, v := range result {
		x[i] = v.Description
	}
	return errorWrapper.WriteJSON(c, http.StatusOK, x)
}
