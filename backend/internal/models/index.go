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

var DataBase *gorm.DB

func ConnectDB() error {
	var err error

	cfg := &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	dsn := getDSNFromEnv()
	db, err := gorm.Open(postgres.Open(dsn), cfg)

	if err != nil {
		return err
	}

	// Only migrate Restaurant table on app startup
	err = db.AutoMigrate(
		&User{},
		&Branch{},
		&Customer{},
		&Inventory{},
		&MenuCategory{},
		&MenuItem{},
		&Notification{},
		&Order{},
		&OrderItem{},
		&Payment{},
		&QRSession{},
		&QRCartItem{}, // Added missing model
		&QRCodeScan{}, // Added missing model
		&Reservation{},
		&Supplier{},
		&Table{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate Restaurant table: %w", err)
	}

	// Assign db to package-level variable
	DataBase = db

	fmt.Println("Database connection established successfully")
	fmt.Println("Restaurant table migrated successfully")
	// Automatically migrate all other tables
	return nil
}
