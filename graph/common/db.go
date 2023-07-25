// @/graph/common/db.go
package common

import (
	"github.com/christiancrawford/go-gqlgen/graph/customTypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDb() (*gorm.DB, error) {
	var err error
	db, err := gorm.Open(postgres.Open("dev.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&customTypes.Todo{})
	return db, nil
}
