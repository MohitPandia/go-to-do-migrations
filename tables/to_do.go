package tables

import (
	"database/sql"
	"time"
)

type Todo struct {
	ID          int            `gorm:"column:todo_id;primaryKey;autoIncrement"`
	UserPID     sql.NullString `gorm:"column:user_pid;not null;type:varchar(40)"`
	CategoryPID sql.NullString `gorm:"column:category_pid;not null;type:varchar(40)"`
	Title       string         `gorm:"column:title;not null;type:varchar(200)"`
	Description string         `gorm:"column:description;type:text"`
	DueDate     time.Time      `gorm:"column:due_date"`
	Completed   bool           `gorm:"column:completed;not null;default:false"`
	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;autoUpdateTime"`

	// Foreign Keys
	// User     Users     `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// Category Categories `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
