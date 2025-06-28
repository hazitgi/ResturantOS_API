package models

import (
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string `gorm:"not null;size:100"`
	Phone         string `gorm:"unique;size:20"`
	Email         string `gorm:"size:255"`
	Address       string `gorm:"type:text"`
	BirthDate     *time.Time
	Anniversary   *time.Time
	TotalOrders   int     `gorm:"default:0"`
	TotalSpent    float64 `gorm:"type:decimal(10,2);default:0"`
	LoyaltyPoints int     `gorm:"default:0"`
	Notes         string  `gorm:"type:text"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
