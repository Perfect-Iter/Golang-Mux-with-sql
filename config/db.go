package config

import (
	"fmt"
	"goMux/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	err    error
	dbHost = os.Getenv("DB_HOST")
	dbUser = "root"
	dbPass = os.Getenv("DB_PASSWORD")
	dbName = "gormdb"
	DNS    = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
)

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	DB.AutoMigrate(&models.User{})
}
