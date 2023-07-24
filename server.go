package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {

	dbUser := goDotEnvVariable("DB_USER")
	dbPassword := goDotEnvVariable("DB_PASSWORD")
	dbName := goDotEnvVariable("DB_NAME")
	dbHost := goDotEnvVariable("DB_HOST")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s", dbUser, dbPassword, dbName, dbHost)

	// Replace the following with your actual PostgreSQL connection details
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database: ", err)
	}

	fmt.Println("Connected to the PostgreSQL database!")

	// Now you can perform database operations using the "db" object.
	// For example, you can query data, insert records, etc.

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = defaultPort
	// }

	// srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)

	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))
}
