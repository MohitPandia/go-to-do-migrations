package tables

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID        int            `gorm:"column:user_id;primaryKey;autoIncrement"`
	PID       sql.NullString `gorm:"column:user_pid;unique;not null;type:varchar(40)"`
	Name      string         `gorm:"column:name;not null;type:varchar(100)"`
	Email     string         `gorm:"column:email;unique;not null;type:varchar(100)"`
	Password  string         `gorm:"column:password;not null;type:varchar(100)"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"` // Adds support for soft deletes
}
