package db

import (
	"database/sql"
	"fmt"
	"os"
)

func ConnectToPostgresDb() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	dbconf := fmt.Sprintf("user=%s password=%s dbname=%s host=%s", dbUser, dbPassword, dbName, dbHost)

	db, err := sql.Open(os.Getenv("DRIVER"), dbconf)
	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil, err
	}
	return db, err

}
