package config

import (
	"os"
)

type Config struct {
	Env string `env:"APP_ENV" default:"development"`
	Port string `env:"PORT" default:"8080"`
}

func Load() *Config {
	cfg := &Config{
		Env: getEnv("ENV", "development"),
		Port: getEnv("PORT", "8080"),
	}
	return cfg
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}