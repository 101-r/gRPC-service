package config

import (
	"fmt"
	"net"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config interface {
	Address() string
	ConnStr() string
}

type config struct {
	ENV            string `yaml:"env" env-default:"local"`
	PostgresConfig `yaml:"postgres"`
	GRPCConfig     `yaml:"grpc"`
}

type PostgresConfig struct {
	POSTGRES_DB       string `yaml:"postgres_db"`
	POSTGRES_USER     string `yaml:"postgres_user"`
	POSTGRES_PASSWORD string `yaml:"postgres_password"`
	POSTGRES_HOST     string `yaml:"postgres_host"`
	POSTGRES_PORT     string `yaml:"postgres_port"`
	SSLMODE           string `yaml:"sslmode"`
}

type GRPCConfig struct {
	GRPC_HOST string `yaml:"grpc_host"`
	GRPC_PORT string `yaml:"grpc_port"`
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return fmt.Errorf("failed to load .env file at path: %s, error: %w", path, err)
	}

	return nil
}

func NewConfig() (Config, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		return nil, fmt.Errorf("config path is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file does not exist at path: %s", configPath)
	}

	var cfg config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, fmt.Errorf("can't read config file at path: %s, error: %w", configPath, err)
	}

	return &cfg, nil
}

func (cfg *config) Address() string {
	return net.JoinHostPort(cfg.GRPC_HOST, cfg.GRPC_PORT)
}

func (cfg *config) ConnStr() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.POSTGRES_USER,
		cfg.POSTGRES_PASSWORD,
		cfg.POSTGRES_HOST,
		cfg.POSTGRES_PORT,
		cfg.POSTGRES_DB,
		cfg.SSLMODE,
	)
}
