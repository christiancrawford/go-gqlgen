package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func ConnectToPostgresDb() (*sql.DB, error) {
	dbDriver := goDotEnvVariable("DB_DRIVER")
	dbUser := goDotEnvVariable("DB_USER")
	dbPassword := goDotEnvVariable("DB_PASSWORD")
	dbName := goDotEnvVariable("DB_NAME")
	dbHost := goDotEnvVariable("DB_HOST")

	dbconf := fmt.Sprintf("user=%s password=%s dbname=%s host=%s", dbUser, dbPassword, dbName, dbHost)

	db, err := sql.Open(dbDriver, dbconf)
	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil, err
	}
	return db, err

}
