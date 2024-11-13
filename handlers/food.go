package handlers

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/phonghaido/healthy-habits/internal/db"
	internal_type "github.com/phonghaido/healthy-habits/internal/types"
	"github.com/phonghaido/healthy-habits/views/components"
	"github.com/phonghaido/healthy-habits/views/pages"
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
	description := c.FormValue("description")
	category := c.FormValue("category")

	requestBody := internal_type.FindFoodReqBody{
		Description: description,
		Category:    category,
	}

	result, err := h.MongoCollection.FindMany(requestBody)
	if err != nil {
		return err
	}

	component := components.SearchResult(result)
	return component.Render(c.Request().Context(), c.Response())
}

func (h *FoodHandler) HandlePOSTFoodDetails(c echo.Context) error {
	var food internal_type.FoundationFood

	if err := json.Unmarshal([]byte(c.FormValue("food")), &food); err != nil {
		return err
	}

	page := pages.FoodDetails(food)

	return page.Render(c.Request().Context(), c.Response())
}
