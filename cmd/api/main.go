package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	apiPort := os.Getenv("PORT")
	if apiPort == "" {
		log.Fatalln("web app port not defined")
	}
	api := echo.New()

	// set middleware
	api.Use(middleware.Logger())
	api.Use(middleware.Recover())

	// Decorate with routes and handlers
	addRoutes(api)

	// Start server
	api.Logger.Fatal(api.Start(":" + apiPort))
}
