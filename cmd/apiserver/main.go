package main

import (
	"log"

	"github.com/JustVlad124/EWallet/internal/app/apiserver"
	"github.com/joho/godotenv"
)

func main() {
	config := apiserver.NewConfig()
	// hardcode
	err := godotenv.Load("./configs/.env")
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
