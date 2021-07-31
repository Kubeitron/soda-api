package main

import "github.com/labstack/echo/v4"

func addRoutes(api *echo.Echo) {
	api.GET("/health", healthcheck).Name = "/health"
}
