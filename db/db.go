package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func ConnectToPostgresDb() {

	dbUser := goDotEnvVariable("DB_USER")
	dbPassword := goDotEnvVariable("DB_PASSWORD")
	dbName := goDotEnvVariable("DB_NAME")
	dbHost := goDotEnvVariable("DB_HOST")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s", dbUser, dbPassword, dbName, dbHost)

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

}
