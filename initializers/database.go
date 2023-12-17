package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB = db
	if err == nil {
		log.Println("Connected to database")
	}
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}
