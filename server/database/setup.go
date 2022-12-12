package database

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func InitDb() (*gorm.DB, error) {
	var err error
	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		log.Fatalf("DSN string is empty")
	}

	dbSql, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Println("Something went wrong with the DSN")
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: dbSql,
	}), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
