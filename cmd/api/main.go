package main

import (
	"log"
	"os"

	"github.com/ducpx/rest-api/config"
	"github.com/ducpx/rest-api/internal/server"
	"github.com/ducpx/rest-api/pkg/logger"
	"github.com/ducpx/rest-api/pkg/utils"
)

func main() {
	log.Println("Starting api server")

	configPath := utils.GetConfigPath(os.Getenv("config"))
	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	cfg, err := config.ParseConfig(cfgFile)

	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode)

	s := server.NewServer(cfg, appLogger)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
