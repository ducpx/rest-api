package server

import (
	"net/http"

	"github.com/ducpx/rest-api/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) MapHandlers(e *echo.Echo) error {

	e.Use(middleware.RequestID())

	v1 := e.Group("/api/v1")

	health := v1.Group("/health")

	health.GET("", func(c echo.Context) error {
		s.logger.Infof("Health check RequestID: %s", utils.GetRequestID(c))
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})
	return nil
}
