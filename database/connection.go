package database

import (
	"database/sql"
	"go-to-do/configs"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, *sql.DB) {
	dsn := "host=" + configs.DB.Host +
		" user=" + configs.DB.Username +
		" password=" + configs.DB.Password +
		" dbname=" + configs.DB.Database +
		" port=" + configs.DB.Port +
		" sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("[Connection], Error in opening db")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("[Connection], Error in setting sqldb")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, sqlDB
}
