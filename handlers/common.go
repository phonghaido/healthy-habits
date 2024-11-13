package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/phonghaido/healthy-habits/internal/db"
	"github.com/phonghaido/healthy-habits/views/pages"
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
	landingPage := pages.LandingPage()
	return landingPage.Render(c.Request().Context(), c.Response())
}
