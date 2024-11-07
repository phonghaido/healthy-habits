package main

import (
	"github.com/labstack/echo/v4"
	"github.com/phonghaido/healthy-habits/handlers"
	"github.com/phonghaido/healthy-habits/internal/config"
	"github.com/phonghaido/healthy-habits/internal/db"
	custom_error "github.com/phonghaido/healthy-habits/pkg/error"
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
	defer db.FoodMongoDBClient.Disconnect()

	db.MealMongoDBClient, err = db.NewMongoDBFMealClient()
	if err != nil {
		logrus.Fatal(err)
	}
	defer db.MealMongoDBClient.Disconnect()

	e := echo.New()

	foodGroup := e.Group("/food")

	foodGroup.POST("", custom_error.ErrorWrapper(handlers.HandleGETFindFood))

	mealGroup := e.Group("/meal")
	mealGroup.POST("", custom_error.ErrorWrapper(handlers.HandlePOSTCreateMealPlan))
	mealGroup.PUT("", custom_error.ErrorWrapper(handlers.HandlePUTUpdateMealPlan))
	mealGroup.DELETE("", custom_error.ErrorWrapper(handlers.HandleDeleteMealPlan))

	e.GET("/", custom_error.ErrorWrapper(handlers.HandleGETLandingPage))

	e.Logger.Fatal(e.Start(":3000"))
}
