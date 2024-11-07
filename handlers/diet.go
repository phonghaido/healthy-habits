package handlers

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	internal_type "github.com/phonghaido/healthy-habits/internal"
	"github.com/phonghaido/healthy-habits/internal/db"
	custom_error "github.com/phonghaido/healthy-habits/pkg/error"
)

func HandlePOSTCreateDietPlan(c echo.Context) error {
	var body internal_type.MealPlan
	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		return err
	}

	if body.Name == "" {
		return custom_error.InvalidRequestBody("name")
	}
	if body.Type == "" {
		return custom_error.InvalidRequestBody("type")
	}

	body.ID = uuid.New().String()

	if err := db.MealMongoDBClient.InsertOne(body); err != nil {
		return err
	}

	return nil
}

func HandlePUTUpdateDietPlan(c echo.Context) error {
	var body internal_type.MealPlan
	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		return err
	}
	if body.ID == "" {
		return custom_error.InvalidRequestBody("id")
	}
	if body.Name == "" {
		return custom_error.InvalidRequestBody("name")
	}
	if body.Type == "" {
		return custom_error.InvalidRequestBody("type")
	}

	if err := db.MealMongoDBClient.UpdateOne(body); err != nil {
		return err
	}
	return nil
}
