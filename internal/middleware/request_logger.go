package middleware

import (
	"time"

	"github.com/ducpx/rest-api/pkg/utils"
	"github.com/labstack/echo/v4"
)

// Request logger middleware
func (mw *MiddlewareManager) RequestLoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)

		req := c.Request()
		res := c.Response()
		status := res.Status
		size := res.Size
		requestID := utils.GetRequestID(c)
		s := time.Since(start).String()

		mw.logger.Infof("RequestID: %s, Method: %s, URI: %s, Status: %v, Size: %v, Time: %s",
			requestID, req.Method, req.URL, status, size, s,
		)
		return err
	}
}
