package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func healthcheck(c echo.Context) error {
	s := &Status{
		Status: "Healthy!",
	}
	content := c.JSON(http.StatusOK, s)
	return content
}
