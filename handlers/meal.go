package handlers

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/phonghaido/healthy-habits/internal/db"
	"github.com/phonghaido/healthy-habits/internal/diet"
	custom_error "github.com/phonghaido/healthy-habits/pkg/error"
)

type DeleteMealReqBody struct {
	IDs []string `json:"ids"`
}

func HandlePOSTCreateMealPlan(c echo.Context) error {
	var body diet.MealPlan
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

func HandlePUTUpdateMealPlan(c echo.Context) error {
	var body diet.MealPlan
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

func HandleDeleteMealPlan(c echo.Context) error {
	var body DeleteMealReqBody
	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		return err
	}

	if len(body.IDs) == 0 {
		return custom_error.InvalidRequestBody("id")
	}

	if len(body.IDs) == 1 {
		if err := db.MealMongoDBClient.DeleteOne(body.IDs[0]); err != nil {
			return err
		}
	} else {
		if err := db.MealMongoDBClient.DeleteMany(body.IDs); err != nil {
			return err
		}
	}
	return nil
}
