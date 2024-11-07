package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	internal_type "github.com/phonghaido/healthy-habits/internal"
	"github.com/phonghaido/healthy-habits/internal/db"
	errorWrapper "github.com/phonghaido/healthy-habits/pkg/error"
)

func HandleGETFindFood(c echo.Context) error {
	var requestBody internal_type.FindFoodReqBody
	if err := json.NewDecoder(c.Request().Body).Decode(&requestBody); err != nil {
		return err
	}

	result, err := db.FoodMongoDBClient.FindMany(requestBody)
	if err != nil {
		return err
	}
	return errorWrapper.WriteJSON(c, http.StatusOK, result)
}
