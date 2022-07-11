package repository

import (
	"context"

	"github.com/ducpx/rest-api/internal/auth"
	"github.com/ducpx/rest-api/internal/models"
	"github.com/ducpx/rest-api/pkg/logger"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
)

type authRepo struct {
	db     *sqlx.DB
	logger logger.Logger
}

// Register implements auth.Repository
func (r *authRepo) Register(ctx context.Context, user *models.User) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authRepo.Register")
	defer span.Finish()

	u := &models.User{}
	if err := r.db.QueryRowxContext(ctx, createUserQuery, &user.FirstName, &user.LastName, &user.Email,
		&user.Password, &user.Role, &user.About, &user.Avatar, &user.PhoneNumber, &user.Address, &user.City,
		&user.Gender, &user.Postcode, &user.Birthday,
	).StructScan(u); err != nil {
		r.logger.Errorf("r.db.QueryRowxContext: %v", err)
		return nil, err
	}
	return u, nil
}

func NewAuthRepo(db *sqlx.DB, logger logger.Logger) auth.Repository {
	return &authRepo{
		db:     db,
		logger: logger,
	}
}
