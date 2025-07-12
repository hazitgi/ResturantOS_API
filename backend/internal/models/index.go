package models

import (
	"fmt"
	"restaurant_os/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getDSNFromConfig(cfg *config.Config) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
}

var DataBase *gorm.DB

func ConnectDB(cfg *config.Config) error {
	var err error
	var db *gorm.DB

	gormCfg := &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent),
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	if cfg.App_Mode == "offline" {
		// use SQLite for offline mode
		db, err = gorm.Open(sqlite.Open("restaurant_os.db"), gormCfg)
		if err != nil {
			return fmt.Errorf("failed to connect to SQLite database: %w", err)
		}
	} else {
		// Use PostgreSQL for non-offline mode
		dsn := getDSNFromConfig(cfg)
		db, err = gorm.Open(postgres.Open(dsn), gormCfg)
		if err != nil {
			return err
		}
	}

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
		&QRCartItem{},
		&QRCodeScan{},
		&Reservation{},
		&Supplier{},
		&Table{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate tables: %w", err)
	}

	DataBase = db

	fmt.Println("Database connection established successfully")
	return nil
}
