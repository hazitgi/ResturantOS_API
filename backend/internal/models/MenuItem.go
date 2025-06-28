package models

import (
	"gorm.io/gorm"
	"time"
)

type MenuItem struct {
	ID           uint   `gorm:"primaryKey"`
	BranchID     uint   `gorm:"not null"`
	Branch       Branch `gorm:"foreignKey:BranchID"`
	CategoryID   *uint
	Category     *MenuCategory `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name         string        `gorm:"not null;size:100"`
	Description  string        `gorm:"type:text"`
	Price        float64       `gorm:"not null;type:decimal(10,2)"`
	CostPrice    float64       `gorm:"type:decimal(10,2)"` // For profit calculation
	Available    bool          `gorm:"default:true"`
	IsVegetarian bool          `gorm:"default:false"`
	IsVegan      bool          `gorm:"default:false"`
	IsGlutenFree bool          `gorm:"default:false"`
	Spiciness    int           `gorm:"default:0"`  // 0-5 scale
	PrepTime     int           `gorm:"default:15"` // minutes
	ImageURL     string        `gorm:"size:500"`
	Ingredients  string        `gorm:"type:text"` // JSON or comma-separated
	Allergens    string        `gorm:"type:text"` // JSON or comma-separated
	SortOrder    int           `gorm:"default:0"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
