package common

import (
	"fmt"
	"log"
	"os"

	"github.com/christiancrawford/go-gqlgen/graph/customTypes"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// This code is loading configuration from a .env file:
func goDotEnvVariable(key string) string {
	// - godotenv.Load(".env") loads the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// - os.Getenv(key) retrieves the value for the given key
	return os.Getenv(key)

	// - The value is returned as a string
}

// This code initializes the database connection:

func InitDb() (*gorm.DB, error) {

	// It loads the database credentials from the .env file using goDotEnvVariable()
	dbUser := goDotEnvVariable("DB_USER")
	dbPassword := goDotEnvVariable("DB_PASSWORD")
	dbName := goDotEnvVariable("DB_NAME")
	dbHost := goDotEnvVariable("DB_HOST")

	// It constructs the database connection string dbconf with the credentials
	dbconf := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", dbHost, dbUser, dbPassword, dbName)

	var err error

	// It opens the database connection using gorm.Open(), passing in the postgres driver and connection string
	db, err := gorm.Open(postgres.Open(dbconf), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	// It runs AutoMigrate to initialize the database schema
	db.AutoMigrate(&customTypes.Todo{})

	// It returns the gorm DB instance to use for database access
	// It configures gorm to use a silent logger
	return db, nil
}
