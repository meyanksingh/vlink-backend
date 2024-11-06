package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() *gorm.DB {
	// Fetch database URL from environment variables
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("Error Loading the ENV File")
	}

	// Setting up GORM logger for detailed logs
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	// Initialize the connection to the PostgreSQL database
	var err error
	var DB *gorm.DB
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	} else {
		fmt.Println("Database connection established successfully!")
	}

	// Optionally, auto-migrate models (replace `YourModel` with actual models)
	// err = DB.AutoMigrate(&models.User{})
	// if err != nil {
	// 	log.Fatalf("Failed to migrate database models: %v", err)
	// }
	return DB
}
