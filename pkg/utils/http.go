package utils

import (
	"context"
	"time"

	"github.com/ducpx/rest-api/pkg/logger"
	"github.com/labstack/echo/v4"
)

// Get RequestId from echo context
func GetRequestID(c echo.Context) string {
	return c.Response().Header().Get(echo.HeaderXRequestID)
}

// ReqIDCtxKey is a key used for the Request ID in context
type ReqIDCtxKey struct{}

// Get request ctx timeout
func GetReqCtxTimeout(c echo.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*15)
	ctx = context.WithValue(ctx, ReqIDCtxKey{}, GetRequestID(c))
	return ctx, cancel
}

// GetReqCtx from echo context
func GetReqCtx(c echo.Context) context.Context {
	return context.WithValue(c.Request().Context(), ReqIDCtxKey{}, GetRequestID(c))
}

// Get config path for local or docker
func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "./config/config-docker"
	}
	return "./config/config-local"
}

// Read request body and validate
func ReadRequest(ctx echo.Context, request interface{}) error {
	if err := ctx.Bind(request); err != nil {
		return err
	}

	return ValidateStructCtx(ctx.Request().Context(), request)
}

// LogResponseError when handler get an error
func LogResponseError(ctx echo.Context, log logger.Logger, err error) {
	log.Errorf(
		"ErrorResponseWithLog, RequestID: %s, Error: %v",
		GetRequestID(ctx), err,
	)
}
