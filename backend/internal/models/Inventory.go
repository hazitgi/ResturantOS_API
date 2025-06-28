package models

import (
    "time"
    "gorm.io/gorm"
)

type Inventory struct {
    ID          uint           `gorm:"primaryKey"`
    BranchID    uint
    ItemName    string
    Unit        string         // e.g. kg, liter, piece
    Quantity    float64
    ReorderLevel float64
    UpdatedAt   time.Time
    CreatedAt   time.Time
    DeletedAt   gorm.DeletedAt `gorm:"index"`
}
