package config

import (
	"errors"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server  ServerConfig
	Metrics Metrics
	Logger  Logger
}

// Server config
type ServerConfig struct {
	AppVersion        string
	Port              string
	Mode              string
	JwtSecretKey      string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	CtxDefaultTimeout time.Duration
	Debug             bool
}

// Metrics config
type Metrics struct {
	URL         string
	ServiceName string
}

// Logger config
type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

// Load config file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

// Parse config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}
	return &c, nil
}
