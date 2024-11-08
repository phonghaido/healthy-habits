package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/phonghaido/healthy-habits/components"
)

type CommonHandler struct {
	Food FoodHandler
	Meal MealHandler
}

func NewCommonHandler(food FoodHandler, meal MealHandler) CommonHandler {
	return CommonHandler{
		Food: food,
		Meal: meal,
	}
}

func (h CommonHandler) HandleGETLandingPage(c echo.Context) error {
	component := components.DefaultPage()
	component.Render(c.Request().Context(), c.Response())
	return nil
}
