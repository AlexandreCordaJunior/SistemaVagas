package config

import (
	//"gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() (*gorm.DB, error) {
	dsn := "host=db user=gorm password=gorm dbname=gorm password=gorm port=5432 sslmode=disable"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//return gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
	//	DisableForeignKeyConstraintWhenMigrating: true,
	//})
}
