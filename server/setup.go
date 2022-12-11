package main

import (
	"github.com/AntonioTrupac/socialHabitsTracker/graph/customTypes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func InitDb() (*gorm.DB, error) {
	var err error

	// create connection to planetscale with gorm
	db, err := gorm.Open(mysql.Open(os.Getenv("DATABASE_URL")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&customTypes.Todo{})

	return db, nil
}
