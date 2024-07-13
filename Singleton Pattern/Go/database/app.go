package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var database *gorm.DB

func ConnectToDatabase() {
	var err error

	dsn := "root:root@tcp(127.0.0.1:3306)/payment"
	database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	mySqlDB, err := database.DB()
	if err != nil {
		fmt.Printf("Error getting *sql.DB from gorm.DB: %v", err)
	}

	mySqlDB.SetMaxOpenConns(1000)
	mySqlDB.SetMaxIdleConns(10)

}

func Migrate() {
	err := database.AutoMigrate(&Payment{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
