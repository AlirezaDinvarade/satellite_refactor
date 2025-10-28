package models

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CheckDatabaseConnection() {
	dsn := os.Getenv("AUTH_DATABASE_URL")
	var err error

	// TODO: use once for open database connection
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("faild to connect database: ", err)
	}
	log.Println("Connectd to database")
}

func CreateTables() {
	err := DB.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("failed to migrate tables:", err)
	}

	log.Println("Tables migrated successfully!")
}
