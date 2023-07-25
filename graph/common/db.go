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
	dbUser := goDotEnvVariable("DB_USER")
	dbPassword := goDotEnvVariable("DB_PASSWORD")
	dbName := goDotEnvVariable("DB_NAME")
	dbHost := goDotEnvVariable("DB_HOST")

	dbconf := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", dbHost, dbUser, dbPassword, dbName)

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
