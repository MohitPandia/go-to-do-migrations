package seeds

import (
	"database/sql"
	"go-to-do/tables"

	"gorm.io/gorm"
)

func Categories(db *gorm.DB) error {
	err := db.Create(&tables.Categories{
		PID:  sql.NullString{String: "cat-001", Valid: true},
		Name: "Work",
	}).Error

	if err != nil {
		return err
	}

	err = db.Create(&tables.Categories{
		PID:  sql.NullString{String: "cat-002", Valid: true},
		Name: "Important",
	}).Error

	if err != nil {
		return err
	}

	err = db.Create(&tables.Categories{
		PID:  sql.NullString{String: "cat-003", Valid: true},
		Name: "Shopping",
	}).Error

	if err != nil {
		return err
	}

	return err
}
