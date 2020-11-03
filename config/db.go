package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetConnection() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
}
