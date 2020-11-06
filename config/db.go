package config

import (
	//"gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() (*gorm.DB, error) {
	dsn := "host=db user=gorm password=gorm dbname=gorm password=gorm port=5432 sslmode=disable"
	var db *gorm.DB
	var err error

	for i := 0; i < 5; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
	}

	if err != nil {
		return nil, err
	}
	return db, err
}
