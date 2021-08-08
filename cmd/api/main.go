package main

import (
	"log"
	"os"

	"github.com/Kubeitron/soda-api/cmd/api/db"
	"github.com/Kubeitron/soda-api/cmd/api/handlers"
	"github.com/Kubeitron/soda-api/cmd/api/store"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
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
	defer func() {
		if err := db.Conn.Disconnect(db.BaseCtx); err != nil {
			panic(err)
		}
	}()

	handler := handlers.New()
	handler.Register(api)

	hch := handlers.NewHealthcheckHandler()
	handler.API.GET("/health", hch.Healthcheck)

	vs := store.NewVegetableStore(db)
	vh := handlers.NewVegetableHandler(vs)
	handler.API.GET("/vegetables", vh.GetVegetables)

	// Start server
	api.Logger.Fatal(api.Start(":" + apiPort))
}
