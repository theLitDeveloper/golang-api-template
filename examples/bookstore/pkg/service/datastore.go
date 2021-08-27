package service

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func initDatastore() {
	var err error
	//
	// Open a connection to a MySQL
	//
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	//
	// Postgres
	//
	// dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%s sslmode=disable TimeZone=Europe/Berlin",
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_NAME"),
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASS"),
	// 	os.Getenv("DB_PORT"))
	// Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	zap.L().Fatal(err.Error())
	// }

	//
	// Auto migrations
	//
	Db.AutoMigrate(&Book{}) // Use your models here
}
