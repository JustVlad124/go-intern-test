package apiserver

import (
	"log"
	"os"

	"github.com/JustVlad124/EWallet/internal/app/store"
	"github.com/joho/godotenv"
)

type Config struct {
	Addr     string
	Port     string
	LogLevel string
	Store    *store.Config
}

func NewConfig() *Config {
	if err := godotenv.Load("./configs/.env"); err != nil {
		log.Fatal(err)
	}
	return &Config{
		Addr:     os.Getenv("ADDR"),
		Port:     os.Getenv("PORT"),
		LogLevel: os.Getenv("LOG_LEVEL"),
		Store:    store.NewConfig(),
	}
}
