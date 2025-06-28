package models

import (
	"gorm.io/gorm"
	"time"
)

type MenuCategory struct {
	ID          uint       `gorm:"primaryKey"`
	BranchID    uint       `gorm:"not null"`
	Branch      Branch     `gorm:"foreignKey:BranchID"`
	Name        string     `gorm:"not null;size:100"`
	Description string     `gorm:"type:text"`
	SortOrder   int        `gorm:"default:0"`
	IsActive    bool       `gorm:"default:true"`
	MenuItems   []MenuItem `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
