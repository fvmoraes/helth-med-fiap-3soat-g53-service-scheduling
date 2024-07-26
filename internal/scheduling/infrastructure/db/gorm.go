package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"helthmed-scheduling/internal/scheduling/domain"
)

var DB *gorm.DB

func Init() {
	var err error
	password := os.Getenv("DATABASE_PASSWORD")
	if password == "" {
		log.Fatal("DATABASE_PASSWORD environment variable is not set")
	}

	dsn := "host=34.44.100.223 user=helthmed password=" + password + " dbname=helthmed-auth port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = DB.AutoMigrate(&domain.Appointment{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}
