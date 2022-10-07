package dto

import (
	"github.com/labstack/echo/v4"
)

type Response struct {
	Status  int       `json:"status,omitempty"`
	Message string    `json:"message,omitempty"`
	Data    *echo.Map `json:"data,omitempty"`
}
