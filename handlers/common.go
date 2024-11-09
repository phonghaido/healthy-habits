package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/phonghaido/healthy-habits/components"
	"github.com/phonghaido/healthy-habits/internal/db"
	internal_type "github.com/phonghaido/healthy-habits/internal/types"
)

type CommonHandler struct {
	Food *db.FoodCollection
	Meal *db.MealCollection
}

func NewCommonHandler(foodH FoodHandler, mealH MealHandler) CommonHandler {
	return CommonHandler{
		Food: foodH.MongoCollection,
		Meal: mealH.MongoCollection,
	}
}

func (h CommonHandler) HandleGETLandingPage(c echo.Context) error {
	body := internal_type.FindFoodReqBody{Category: "", Description: ""}
	food, err := h.Food.FindMany(body)
	if err != nil {
		return err
	}
	component := components.DefaultPage(food)
	component.Render(c.Request().Context(), c.Response())
	return nil
}
