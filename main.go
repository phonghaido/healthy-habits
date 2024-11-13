package main

import (
	"fmt"

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

	commonConfig := config.GetCommonConfig()

	mongoClient, err := db.NewMongoClient()
	if err != nil {
		logrus.Fatal(err)
	}
	defer mongoClient.Disconnect()

	foodHandler := handlers.NewFoodHandler(mongoClient)
	mealHandler := handlers.NewMealHandler(mongoClient)
	commonHandler := handlers.NewCommonHandler(foodHandler, mealHandler)

	e := echo.New()
	e.Static("/assets", "assets")
	e.GET("/", custom_error.ErrorWrapper(commonHandler.HandleGETLandingPage))

	foodGroup := e.Group("/food")
	foodGroup.POST("", custom_error.ErrorWrapper(foodHandler.HandleGETFindFood))
	foodGroup.GET("/details", custom_error.ErrorWrapper(foodHandler.HandlePOSTFoodDetails))

	mealGroup := e.Group("/meal")
	mealGroup.POST("", custom_error.ErrorWrapper(mealHandler.HandlePOSTCreateMealPlan))
	mealGroup.PUT("", custom_error.ErrorWrapper(mealHandler.HandlePUTUpdateMealPlan))
	mealGroup.DELETE("", custom_error.ErrorWrapper(mealHandler.HandleDeleteMealPlan))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", commonConfig.Port)))
}
