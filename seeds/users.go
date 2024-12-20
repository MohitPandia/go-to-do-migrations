package seeds

import (
	"database/sql"
	"go-to-do/tables"
	"time"

	"gorm.io/gorm"
)

func Users(db *gorm.DB) error {
	err := db.Create(&tables.Users{
		PID:       sql.NullString{String: "user-001", Valid: true},
		Name:      "Yash Pandia",
		Email:     "yash@example.com",
		Password:  "hashed_password_1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).Error

	if err != nil {
		return err
	}

	err = db.Create(&tables.Users{
		PID:       sql.NullString{String: "user-002", Valid: true},
		Name:      "Mohit Pandia",
		Email:     "mohit@example.com",
		Password:  "hashed_password_2",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).Error

	if err != nil {
		return err
	}

	return err
}
