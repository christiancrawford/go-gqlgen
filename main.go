package main

import (
	"log"

	"github.com/christiancrawford/go-gqlgen/server"
	"github.com/joho/godotenv"
)

func envLoad() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading env target")
	}
}

func main() {
	envLoad() // Load from .env file
	server.StartServer()
}
