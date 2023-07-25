// @/graph/common/db.go
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

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func InitDb() (*gorm.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	dbconf := fmt.Sprintf("user=%s password=%s dbname=%s host=%s", dbUser, dbPassword, dbName, dbHost)

	var err error
	db, err := gorm.Open(postgres.Open(dbconf), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&customTypes.Todo{})
	return db, nil
}
