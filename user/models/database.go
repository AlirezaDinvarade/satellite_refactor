package models

import (
	"log"
	"os"
	"satellite/user/types"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Check_database() {
	dsn := os.Getenv("AUTH_DATABASE_URL")
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("faild to connect database: ", err)
	}
	log.Println("Connectd to database")
}

func CreateTables() {
	err := DB.AutoMigrate(&types.User{})
	if err != nil {
		log.Fatal("failed to migrate tables:", err)
	}

	log.Println("Tables migrated successfully!")
}
