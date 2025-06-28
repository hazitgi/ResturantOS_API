package models

import (
    "time"
    "gorm.io/gorm"
)

type Restaurant struct {
    ID          uint           `gorm:"primaryKey"`
    Name        string         `gorm:"not null;size:100"`
    Email       string         `gorm:"unique;size:255"`
    Phone       string         `gorm:"size:20"`
    Address     string         `gorm:"type:text"`
    Website     string         `gorm:"size:255"`
    TaxNumber   string         `gorm:"size:50"`
    Currency    string         `gorm:"size:3;default:'USD'"` // ISO currency code
    TimeZone    string         `gorm:"size:50;default:'UTC'"`
    IsActive    bool           `gorm:"default:true"`
    Branches    []Branch
    Users       []User
    CreatedAt   time.Time
    UpdatedAt   time.Time
    DeletedAt   gorm.DeletedAt `gorm:"index"`
}