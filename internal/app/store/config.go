package store

import (
	"fmt"
	"os"
)

type Config struct {
	DatabaseURL string
}

func NewConfig() *Config {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=%s timezone=UTC connect_timeout=5",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_SSL_MODE"))

	if os.Getenv("DATABASE_PASS") != "" {
		dsn = fmt.Sprintf("%s password=%s", dsn, os.Getenv("DATABASE_PASS"))
	}

	return &Config{
		DatabaseURL: dsn,
	}
}
