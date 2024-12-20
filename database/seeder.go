package database

import (
	"go-to-do/seeds"

	"gorm.io/gorm"
)

type seed struct {
	TableName string
	Run       func(*gorm.DB) error
}

// here we need to modify seed functions
func Seeder(db *gorm.DB) []seed {
	users := seed{TableName: "users",
		Run: func(d *gorm.DB) error { return seeds.Users(db) }}
	todo := seed{TableName: "todo",
		Run: func(d *gorm.DB) error { return seeds.Todo(db) }}
	categories := seed{TableName: "categories",
		Run: func(d *gorm.DB) error { return seeds.Categories(db) }}

	return []seed{
		users,
		todo,
		categories,
	}
}
