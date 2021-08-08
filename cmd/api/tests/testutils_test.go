package tests

import (
	"log"
	"os"

	"github.com/Kubeitron/soda-api/cmd/api/db"
	"github.com/Kubeitron/soda-api/cmd/api/handlers"
	"github.com/Kubeitron/soda-api/cmd/api/store"
	_ "github.com/Kubeitron/soda-api/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title SODA API
// @version 0.1
// @description Contains data on Kubernetes resources

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
func MockAPI() (h *handlers.Handler) {
	apiPort := os.Getenv("PORT")
	if apiPort == "" {
		log.Fatalln("API port not defined")
	}
	api := echo.New()

	// set middleware
	api.Use(middleware.Logger())
	api.Use(middleware.Recover())

	// Establish db client
	db := db.NewMongodb()

	handler := handlers.New()
	handler.Register(api)

	hch := handlers.NewHealthcheckHandler()
	handler.API.GET("/health", hch.Healthcheck)

	handler.API.GET("/swagger/*", echoSwagger.WrapHandler)

	vs := store.NewVegetableStore(db)
	vh := handlers.NewVegetableHandler(vs)
	handler.API.GET("/vegetables", vh.GetVegetables)

	return handler
}
