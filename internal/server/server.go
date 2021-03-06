package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ducpx/rest-api/config"
	"github.com/ducpx/rest-api/pkg/logger"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"

	"github.com/labstack/echo/v4"
)

const (
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5
)

type Server struct {
	echo        *echo.Echo
	cfg         *config.Config
	db          *sqlx.DB
	redisClient *redis.Client
	logger      logger.Logger
}

func NewServer(cfg *config.Config, db *sqlx.DB, redisClient *redis.Client, logger logger.Logger) *Server {
	return &Server{
		echo:        echo.New(),
		cfg:         cfg,
		db:          db,
		redisClient: redisClient,
		logger:      logger,
	}
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr:           s.cfg.Server.Port,
		ReadTimeout:    time.Second * s.cfg.Server.ReadTimeout,
		WriteTimeout:   time.Second * s.cfg.Server.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		s.logger.Infof("Server is listening on PORT: %s", s.cfg.Server.Port)
		s.echo.HideBanner = true
		if err := s.echo.StartServer(server); err != nil {
			s.logger.Fatalf("Error starting Server: ", err)
		}
	}()

	if err := s.MapHandlers(s.echo); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	return s.echo.Shutdown(ctx)
}
