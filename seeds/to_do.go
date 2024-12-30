package seeds

import (
	"database/sql"
	"go-to-do/tables"
	"time"

	"gorm.io/gorm"
)

func Todo(db *gorm.DB) error {
	err := db.Create(&tables.Todo{
		PID:         sql.NullString{String: "todo-001", Valid: true},
		UserPID:     sql.NullString{String: "user-001", Valid: true},
		CategoryPID: sql.NullString{String: "cat-001", Valid: true},
		Title:       "Complete Project Documentation",
		Description: "Write and finalize all project-related documentation.",
		DueDate:     time.Now().AddDate(0, 0, 7), // Due in 7 days
		Completed:   false,
	}).Error

	if err != nil {
		return err
	}

	err = db.Create(&tables.Todo{
		PID:         sql.NullString{String: "todo-002", Valid: true},
		UserPID:     sql.NullString{String: "user-001", Valid: true},
		CategoryPID: sql.NullString{String: "cat-002", Valid: true},
		Title:       "Go for Health Checkup",
		Description: "go bro go",
		DueDate:     time.Now().AddDate(0, 0, 5), // Due in 5 days
		Completed:   false,
		DeletedAt:   gorm.DeletedAt{Time: time.Now(), Valid: true}, // Mark as soft-deleted
	}).Error

	err = db.Create(&tables.Todo{
		PID:         sql.NullString{String: "todo-003", Valid: true},
		UserPID:     sql.NullString{String: "user-002", Valid: true},
		CategoryPID: sql.NullString{String: "cat-001", Valid: true},
		Title:       "Go for Health Checkup",
		Description: "go bro go",
		DueDate:     time.Now().AddDate(0, 0, 4),
		Completed:   false,
		DeletedAt:   gorm.DeletedAt{Time: time.Now(), Valid: true}, // Mark as soft-deleted
	}).Error

	if err != nil {
		return err
	}

	return err
}
