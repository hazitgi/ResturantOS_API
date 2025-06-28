package models

import (
    "time"
    "gorm.io/gorm"
)

type TableStatus string

const (
    TableAvailable TableStatus = "AVAILABLE"
    TableOccupied  TableStatus = "OCCUPIED"
    TableReserved  TableStatus = "RESERVED"
)

type Table struct {
    ID        uint           `gorm:"primaryKey"`
    Number    string         `gorm:"not null"`
    BranchID  uint
    Status    TableStatus    `gorm:"type:VARCHAR(20);default:'AVAILABLE'"`
    Orders    []Order
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}
