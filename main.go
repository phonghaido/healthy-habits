package main

import (
	"github.com/labstack/echo/v4"
	"github.com/phonghaido/healthy-habits/handlers"
	"github.com/phonghaido/healthy-habits/internal/config"
	"github.com/phonghaido/healthy-habits/internal/db"
	errorWrapper "github.com/phonghaido/healthy-habits/pkg/error"
	"github.com/sirupsen/logrus"
)

func main() {
	err := config.SetupViper()
	if err != nil {
		logrus.Fatal(err)
	}

	db.FoodMongoDBClient, err = db.NewMongoDBFoodClient()
	if err != nil {
		logrus.Fatal(err)
	}

	e := echo.New()

	nutGroup := e.Group("/nutrients")

	nutGroup.GET("", errorWrapper.ErrorWrapper(handlers.HandleGETAllNutrient))

	e.GET("/", errorWrapper.ErrorWrapper(handlers.HandleGETLandingPage))

	e.Logger.Fatal(e.Start(":3000"))
}
