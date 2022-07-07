package http

import (
	"github.com/ducpx/rest-api/internal/auth"
	"github.com/labstack/echo/v4"
)

func MapAuthRoutes(authGroup *echo.Group, h auth.Handlers) {
	authGroup.POST("/register", h.Register())
}
