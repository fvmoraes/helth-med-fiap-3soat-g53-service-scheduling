package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"helthmed-scheduling/internal/scheduling/domain"
)

var DB *gorm.DB

func Init() {
	var err error
	dsn := "host=localhost user=youruser password=yourpassword dbname=helthmed_scheduling port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = DB.AutoMigrate(&domain.Appointment{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}