package handlers

import (
	"net/http"

	"github.com/Kubeitron/soda-api/cmd/api/resources"
	"github.com/labstack/echo/v4"
)

type (
	HealthcheckHandler struct {
		// does not require a store
	}
)

func NewHealthcheckHandler() (h *HealthcheckHandler) {
	return &HealthcheckHandler{}
}

func (h *HealthcheckHandler) Healthcheck(c echo.Context) error {
	s := resources.Status{
		Status: "Healthy!",
	}
	content := c.JSON(http.StatusOK, s)
	return content
}
