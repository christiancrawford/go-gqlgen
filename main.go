package main

import (
	"github.com/christiancrawford/go-gqlgen/db"
	"github.com/christiancrawford/go-gqlgen/server"
)

func main() {
	db.ConnectToPostgresDb()
	server.StartServer()
}
