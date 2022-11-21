package database

import (
	"tubes/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	//todo
	// dsn

	dsn := "root:@tcp(127.0.0.1:3306)/pam?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//open database
	db = db.Debug()
	autoMigrate()
}

func GetDb() *gorm.DB {
	return db
}

func autoMigrate() {
	db.AutoMigrate(&models.Order{}, &models.Item{}, &models.User{}, &models.Product{})
}
