package models

import (
    "time"
    "gorm.io/gorm"
)

type OrderStatus string

const (
    OrderPending    OrderStatus = "PENDING"
    OrderPreparing  OrderStatus = "PREPARING"
    OrderReady      OrderStatus = "READY"
    OrderCompleted  OrderStatus = "COMPLETED"
    OrderCancelled  OrderStatus = "CANCELLED"
)

type Order struct {
    ID         uint           `gorm:"primaryKey"`
    TableID    *uint
    BranchID   uint
    UserID     *uint           // Waiter/Cashier who placed
    Total      float64
    Status     OrderStatus     `gorm:"type:VARCHAR(20);default:'PENDING'"`
    OrderItems []OrderItem
    CreatedAt  time.Time
    UpdatedAt  time.Time
    DeletedAt  gorm.DeletedAt `gorm:"index"`
}
