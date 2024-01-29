package apiserver

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Addr        string
	Port        string
	LogLevel    string
	DatabaseURL string
}

func NewConfig() *Config {
	if err := godotenv.Load("./configs/.env"); err != nil {
		log.Fatal(err)
	}
	return &Config{
		Addr:        os.Getenv("ADDR"),
		Port:        os.Getenv("PORT"),
		LogLevel:    os.Getenv("LOG_LEVEL"),
		DatabaseURL: preapareDatabaseURL(),
	}
}

func preapareDatabaseURL() string {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=%s timezone=UTC connect_timeout=5",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_SSL_MODE"),
	)

	if os.Getenv("DATABASE_PASS") != "" {
		dsn = fmt.Sprintf("%s password=%s", dsn, os.Getenv("DATABASE_PASS"))
	}

	return dsn
}
