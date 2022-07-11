package usecase

import (
	"context"

	"github.com/ducpx/rest-api/config"
	"github.com/ducpx/rest-api/internal/auth"
	"github.com/ducpx/rest-api/internal/models"
	"github.com/ducpx/rest-api/pkg/logger"
	"github.com/opentracing/opentracing-go"
)

type authUC struct {
	cfg      *config.Config
	authRepo auth.Repository
	logger   logger.Logger
}

// Register implements auth.Usecase
func (uc *authUC) Register(ctx context.Context, user *models.User) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authUC.Register")
	defer span.Finish()

	userCreated, err := uc.authRepo.Register(ctx, user)
	if err != nil {
		return nil, err
	}
	return userCreated, nil
}

func NewAuthUC(cfg *config.Config, authRepo auth.Repository, logger logger.Logger) auth.Usecase {
	return &authUC{
		cfg:      cfg,
		authRepo: authRepo,
		logger:   logger,
	}
}
