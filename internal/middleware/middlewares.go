package middleware

import (
	"github.com/ducpx/rest-api/config"
	"github.com/ducpx/rest-api/pkg/logger"
)

// Middleware manager
type MiddlewareManager struct {
	cfg    *config.Config
	logger logger.Logger
}

// Middleware manager constructor
func NewMiddlewareManager(cfg *config.Config, logger logger.Logger) *MiddlewareManager {
	return &MiddlewareManager{
		cfg:    cfg,
		logger: logger,
	}
}
