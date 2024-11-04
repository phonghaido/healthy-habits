package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	errorWrapper "github.com/phonghaido/healthy-habits/pkg/error"
)

func HandleGETAllNutrient(c echo.Context) error {
	return errorWrapper.WriteJSON(c, http.StatusOK, map[string]string{"msg": "All nutrients"})
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
