// This main.go file is the entry point for the application. It:
package main

// Imports the server package
import "github.com/christiancrawford/go-gqlgen/server"

// Defines a main() function
func main() {
	// Calls server.StartServer()
	// Starts the GraphQL server
	server.StartServer()
}

// This allows the server code to be encapsulated in the server package,
// while main.go handles starting it up.
