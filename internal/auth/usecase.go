package auth

import (
	"context"

	"github.com/ducpx/rest-api/internal/models"
)

type Usecase interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
}
