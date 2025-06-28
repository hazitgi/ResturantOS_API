package models

import (
	"gorm.io/gorm"
	"time"
)

type Supplier struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;size:100"`
	Contact   string `gorm:"size:100"`
	Phone     string `gorm:"size:20"`
	Email     string `gorm:"size:255"`
	Address   string `gorm:"type:text"`
	IsActive  bool   `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
