package models

import (
	"gorm.io/gorm"
	"time"
)

type InventoryUnit string

const (
	UnitKG     InventoryUnit = "KG"
	UnitGram   InventoryUnit = "GRAM"
	UnitLiter  InventoryUnit = "LITER"
	UnitML     InventoryUnit = "ML"
	UnitPiece  InventoryUnit = "PIECE"
	UnitPack   InventoryUnit = "PACK"
	UnitBottle InventoryUnit = "BOTTLE"
)

type Inventory struct {
	ID           uint          `gorm:"primaryKey"`
	BranchID     uint          `gorm:"not null"`
	Branch       Branch        `gorm:"foreignKey:BranchID"`
	ItemName     string        `gorm:"not null;size:100"`
	ItemCode     string        `gorm:"size:50"` // SKU or item code
	Category     string        `gorm:"size:50"`
	Unit         InventoryUnit `gorm:"type:VARCHAR(20);not null"`
	CurrentStock float64       `gorm:"type:decimal(10,3);default:0"`
	ReorderLevel float64       `gorm:"type:decimal(10,3);default:0"`
	MaxLevel     float64       `gorm:"type:decimal(10,3);default:0"`
	UnitCost     float64       `gorm:"type:decimal(10,2);default:0"`
	SupplierName string        `gorm:"size:100"`
	LastOrdered  *time.Time
	ExpiryDate   *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
