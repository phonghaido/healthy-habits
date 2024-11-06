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

	foodGroup := e.Group("/food")

	foodGroup.GET("/list", errorWrapper.ErrorWrapper(handlers.HandleGETFindFood))
	foodGroup.POST("/list/category", errorWrapper.ErrorWrapper(handlers.HandlePOSTFindFoodByCategory))
	foodGroup.POST("/list/description", errorWrapper.ErrorWrapper(handlers.HandlePOSTFindFoodByDescription))

	e.GET("/", errorWrapper.ErrorWrapper(handlers.HandleGETLandingPage))

	e.Logger.Fatal(e.Start(":3000"))
}
