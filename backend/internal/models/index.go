package models

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getDSNFromEnv() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
}

var DB *gorm.DB

func ConnectDB() error {

	var err error

	cfg := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}

	dsn := getDSNFromEnv()
	db, err := gorm.Open(postgres.Open(dsn), cfg)

	if err != nil {
		return err
	}

	// Run AutoMigrate for all models
	db.AutoMigrate(
		&User{},
		&Restaurant{},
		&Branch{},
		&Table{},
		&MenuItem{},
		&Order{},
		&OrderItem{},
		&Inventory{},
	)

	// Optionally, assign db to a package-level variable if needed
	DB = db

	fmt.Println("Database connection established successfully")

	return nil
}
