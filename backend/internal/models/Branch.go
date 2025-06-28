package models

import (
    "time"
    "gorm.io/gorm"
)

type Branch struct {
    ID            uint           `gorm:"primaryKey"`
    RestaurantID  uint           `gorm:"not null"`
    Restaurant    Restaurant     `gorm:"foreignKey:RestaurantID"`
    Name          string         `gorm:"not null;size:100"`
    Location      string         `gorm:"type:text"`
    Phone         string         `gorm:"size:20"`
    Email         string         `gorm:"size:255"`
    ManagerID     *uint          // Branch manager
    Manager       *User          `gorm:"foreignKey:ManagerID"`
    IsActive      bool           `gorm:"default:true"`
    OpeningHours  string         `gorm:"type:json"` // JSON for flexible hours
    Tables        []Table
    MenuItems     []MenuItem
    Orders        []Order
    Inventory     []Inventory
    Users         []User
    CreatedAt     time.Time
    UpdatedAt     time.Time
    DeletedAt     gorm.DeletedAt `gorm:"index"`
}
