package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

type (
	Config struct {
		Env  string `yaml:"env" env-required:"true"`
		HTTP `yaml:"http" env-required:"true"`
		PG   `yaml:"postgres"`
	}

	HTTP struct {
		Address     string        `yaml:"address"       env-default:"localhost:8080"`
		Timeout     time.Duration `yaml:"timeout"       env-default:"4s"`
		IdleTimeout time.Duration `yaml:"idle_timeout"  env-default:"60s"`
	}

	PG struct {
		PoolMax  int    `env-required:"true"     yaml:"pool_max"   env:"PG_POOL_MAX"`
		User     string `env-required:"true"     yaml:"user"       env:"PG_USER"`
		Password string `env-required:"true"     yaml:"password"   env:"PG_PASSWORD"`
		Host     string `env-required:"true"     yaml:"host"       env:"PG_HOST"`
		Port     string `env-required:"true"     yaml:"port"       env:"PG_PORT"`
		Name     string `env-required:"true"     yaml:"name"       env:"PG_NAME"`
	}
)

const configPath = "./config/config.yml"

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(configPath, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
