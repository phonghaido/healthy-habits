package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	errorWrapper "github.com/phonghaido/healthy-habits/pkg/error"
)

func HandleGETLandingPage(c echo.Context) error {
	return errorWrapper.WriteJSON(c, http.StatusOK, map[string]string{"msg": "this is landing page"})
}
