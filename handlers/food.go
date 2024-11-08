package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/phonghaido/healthy-habits/internal/db"
	internal_type "github.com/phonghaido/healthy-habits/internal/types"
	errorWrapper "github.com/phonghaido/healthy-habits/pkg/error"
)

type FoodHandler struct {
	MongoCollection *db.FoodCollection
}

func NewFoodHandler(c *db.MongoClient) FoodHandler {
	coll := db.NewFoodCollection(c)
	return FoodHandler{
		MongoCollection: coll,
	}
}

func (h *FoodHandler) HandleGETFindFood(c echo.Context) error {
	var requestBody internal_type.FindFoodReqBody
	if err := json.NewDecoder(c.Request().Body).Decode(&requestBody); err != nil {
		return err
	}

	result, err := h.MongoCollection.FindMany(requestBody)
	if err != nil {
		return err
	}
	return errorWrapper.WriteJSON(c, http.StatusOK, result)
}
