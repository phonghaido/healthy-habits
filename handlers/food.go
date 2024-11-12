package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/phonghaido/healthy-habits/internal/db"
	internal_type "github.com/phonghaido/healthy-habits/internal/types"
	"github.com/phonghaido/healthy-habits/views/components"
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

	fmt.Println(requestBody)
	result, err := h.MongoCollection.FindMany(requestBody)
	if err != nil {
		return err
	}

	html := components.SearchResult(result)
	return html.Render(c.Request().Context(), c.Response())
}
