package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	DBPath         string
	AllowedOrigins []string
}

func Load() *Config {
	_ = godotenv.Load()
	return &Config{
		Port:           getEnv("PORT", "9136"),
		DBPath:         getEnv("DB_PATH", "data/db.sqlite"),
		AllowedOrigins: strings.Split(getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:3000"), ","),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
