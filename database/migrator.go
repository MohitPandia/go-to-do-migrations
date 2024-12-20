package database

import (
	"go-to-do/tables"

	"gorm.io/gorm"
)

type Migrate struct {
	TableName string
	Run       func(*gorm.DB) error
}

func AutoMigrate(db *gorm.DB) []Migrate {

	var users tables.Users
	var categories tables.Categories

	usersM := Migrate{TableName: "users",
		Run: func(d *gorm.DB) error { return db.AutoMigrate(&users) },
	}
	todoM := Migrate{TableName: "todo",
		Run: func(d *gorm.DB) error { return db.AutoMigrate(&tables.Todo{}) },
	}
	categoriesM := Migrate{TableName: "categories",
		Run: func(d *gorm.DB) error { return db.AutoMigrate(&categories) },
	}

	return []Migrate{
		usersM,
		todoM,
		categoriesM,
	}
}
