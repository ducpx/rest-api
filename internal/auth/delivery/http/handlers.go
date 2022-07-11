package http

import (
	"net/http"

	"github.com/ducpx/rest-api/config"
	"github.com/ducpx/rest-api/internal/auth"
	"github.com/ducpx/rest-api/internal/models"
	"github.com/ducpx/rest-api/pkg/logger"
	"github.com/ducpx/rest-api/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
)

type authHandlers struct {
	cfg    *config.Config
	authUc auth.Usecase
	logger logger.Logger
}

// Register implements auth.Handlers
// @Summary Register a new user
// @Description Register a new user, returns user and token
// @Tags auth
// @Accept json
// @Produce json
// @Success 201 {object} models.User
// @Router /auth/register [post]
func (h *authHandlers) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetReqCtx(c), "auth.register")
		defer span.Finish()

		user := &models.User{}

		if err := utils.ReadRequest(c, user); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		userCreated, err := h.authUc.Register(ctx, user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusCreated, map[string]*models.User{"user": userCreated})
	}
}

func NewAuthHandlers(cfg *config.Config, authUc auth.Usecase, logger logger.Logger) auth.Handlers {
	return &authHandlers{
		cfg:    cfg,
		authUc: authUc,
		logger: logger,
	}
}
