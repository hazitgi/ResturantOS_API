package models

import (
	"time"
)

type OrderItemStatus string

const (
	OrderItemPending   OrderItemStatus = "PENDING"
	OrderItemPreparing OrderItemStatus = "PREPARING"
	OrderItemReady     OrderItemStatus = "READY"
	OrderItemServed    OrderItemStatus = "SERVED"
	OrderItemCancelled OrderItemStatus = "CANCELLED"
)

type OrderItem struct {
	ID         uint            `gorm:"primaryKey"`
	OrderID    uint            `gorm:"not null"`
	Order      Order           `gorm:"foreignKey:OrderID"`
	MenuItemID uint            `gorm:"not null"`
	MenuItem   MenuItem        `gorm:"foreignKey:MenuItemID"`
	Quantity   int             `gorm:"not null;default:1"`
	UnitPrice  float64         `gorm:"type:decimal(10,2);not null"`
	TotalPrice float64         `gorm:"type:decimal(10,2);not null"`
	Status     OrderItemStatus `gorm:"type:VARCHAR(20);default:'PENDING'"`
	Notes      string          `gorm:"type:text"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
