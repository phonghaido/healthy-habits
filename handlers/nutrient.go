package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/phonghaido/healthy-habits/internal/db"
	errorWrapper "github.com/phonghaido/healthy-habits/pkg/error"
	"go.mongodb.org/mongo-driver/bson"
)

func HandleGETAllNutrient(c echo.Context) error {
	result, err := db.FoodMongoDBClient.FindMany(bson.D{{}})
	if err != nil {
		return err
	}
	return errorWrapper.WriteJSON(c, http.StatusOK, result)
}

func HandleGETNutrientByType(c echo.Context) error {
	foodType := c.Param("type")

	return errorWrapper.WriteJSON(c, http.StatusOK, map[string]string{"msg": fmt.Sprintf("food type: %s", foodType)})
}

func HandleGETNutrientDetail(c echo.Context) error {
	foodType := c.Param("type")
	foodName := c.Param("name")

	return errorWrapper.WriteJSON(c, http.StatusOK, map[string]string{"msg": fmt.Sprintf("food type: %s, food name: %s", foodType, foodName)})
}
