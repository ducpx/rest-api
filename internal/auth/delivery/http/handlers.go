package http

import (
	"net/http"

	"github.com/ducpx/rest-api/config"
	"github.com/ducpx/rest-api/internal/auth"
	"github.com/ducpx/rest-api/pkg/logger"
	"github.com/ducpx/rest-api/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
)

type authHandlers struct {
	cfg    *config.Config
	logger logger.Logger
}

// Register implements auth.Handlers
func (h *authHandlers) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetReqCtx(c), "auth.register")
		defer span.Finish()

		reqID := ctx.Value(utils.ReqIDCtxKey{}).(string)
		h.logger.Infof("Request ID in register: %s", reqID)
		return c.JSON(http.StatusOK, map[string]string{"register": reqID})
	}
}

func NewAuthHandlers(cfg *config.Config, logger logger.Logger) auth.Handlers {
	return &authHandlers{
		cfg:    cfg,
		logger: logger,
	}
}
