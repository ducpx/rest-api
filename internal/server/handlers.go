package server

import (
	"net/http"

	appMiddleware "github.com/ducpx/rest-api/internal/middleware"
	"github.com/ducpx/rest-api/pkg/metric"
	"github.com/ducpx/rest-api/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	metr, err := metric.CreateMetrics(s.cfg.Metrics.URL, s.cfg.Metrics.ServiceName)
	if err != nil {
		s.logger.Errorf("CreateMetrics error: %v", err)
	}
	s.logger.Infof(
		"Metrics Available URL: %s, ServiceName: %s",
		s.cfg.Metrics.URL, s.cfg.Metrics.ServiceName,
	)
	mw := appMiddleware.NewMiddlewareManager(s.cfg, s.logger)
	e.Use(middleware.RequestID())
	e.Use(mw.MetricsMiddleware(metr))

	v1 := e.Group("/api/v1")

	health := v1.Group("/health")

	health.GET("", func(c echo.Context) error {
		s.logger.Infof("Health check RequestID: %s", utils.GetRequestID(c))
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})
	return nil
}
