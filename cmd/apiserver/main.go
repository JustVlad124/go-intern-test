package main

import (
	"log"

	"github.com/JustVlad124/EWallet/internal/app/apiserver"
	"github.com/joho/godotenv"
)

func main() {
	config := apiserver.NewConfig()
	// hardcoded
	err := godotenv.Load("./configs/.env")
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
