package database

import (
	"graphql/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=Ashwini dbname=graph port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	err = db.Migrator().AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}
	err = db.Migrator().AutoMigrate(&models.Company{}, &models.Job{})
	if err != nil {
		return nil, err
	}
	return db, nil

}
