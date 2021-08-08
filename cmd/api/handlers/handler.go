package handlers

import (
	"github.com/labstack/echo/v4"
)

type (
	Handler struct {
		API *echo.Echo
	}
)

func New() (h *Handler) {
	return &Handler{
		nil,
	}
}

func (h *Handler) Register(api *echo.Echo) {
	h.API = api
}
