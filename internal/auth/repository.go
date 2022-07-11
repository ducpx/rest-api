package auth

import (
	"context"

	"github.com/ducpx/rest-api/internal/models"
)

// Postgres interface
type Repository interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
}

// Cache interface
type RedisRepository interface {
}
