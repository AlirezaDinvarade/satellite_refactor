package stores

import (
	"log"
	"os"
	"satellite/user/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB


type Store struct {
	User UserStore
}

func CheckDatabaseConnection() {
	dsn := os.Getenv("AUTH_DATABASE_URL")
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("faild to connect database: ", err)
	}
	log.Println("Connectd to database")
}

func CreateTables() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("failed to migrate tables:", err)
	}

	log.Println("Tables migrated successfully!")
}
