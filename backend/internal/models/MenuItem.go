package models

import (
    "gorm.io/gorm"
    "time"
)

type MenuItem struct {
    ID          uint           `gorm:"primaryKey"`
    BranchID    uint
    Name        string         `gorm:"not null"`
    Description string
    Price       float64        `gorm:"not null"`
    Available   bool           `gorm:"default:true"`
    Category    string
    CreatedAt   time.Time
    UpdatedAt   time.Time
    DeletedAt   gorm.DeletedAt `gorm:"index"`
}
