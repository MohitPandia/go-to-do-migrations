package tables

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Categories struct {
	ID        int            `gorm:"column:category_id;primaryKey;autoIncrement"`
	PID       sql.NullString `gorm:"column:category_pid;unique;not null;type:varchar(40)"`
	Name      string         `gorm:"column:name;not null;unique;type:varchar(100)"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"` // Adds support for soft deletes
}
