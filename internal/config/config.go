package config

import (
	"strings"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

	"project/pkg/logger"
)

type Config struct {
	ServiceName  string       `envconfig:"SERVICE_NAME" validate:"required"`
	Host         string       `envconfig:"HOST" validate:"required"`
	Port         string       `envconfig:"PORT" validate:"required,startswith=:"`
	LogLevel     logger.Level `envconfig:"LOG_LEVEL"`
	WorkersCount uint         `envconfig:"WORKERS_COUNT"`
}

func Read() (*Config, error) {
	validate := validator.New()
	// Omit error as in case of non-existing .env.local it fill respond with error
	_ = godotenv.Overload("./cmd/.env", "./cmd/.env.local")

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	err := validate.Struct(cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c Config) HostWithoutProtocol() string {
	wo := strings.TrimPrefix(c.Host, "http://")
	return strings.TrimPrefix(wo, "https://")
}
