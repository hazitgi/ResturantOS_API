package models

import (
	"time"
)

type PaymentMethod string

const (
	PaymentCash       PaymentMethod = "CASH"
	PaymentCard       PaymentMethod = "CARD"
	PaymentUPI        PaymentMethod = "UPI"
	PaymentWallet     PaymentMethod = "WALLET"
	PaymentNetBanking PaymentMethod = "NET_BANKING"
)

type Payment struct {
	ID              uint          `gorm:"primaryKey"`
	OrderID         uint          `gorm:"not null"`
	Order           Order         `gorm:"foreignKey:OrderID"`
	Amount          float64       `gorm:"type:decimal(10,2);not null"`
	Method          PaymentMethod `gorm:"type:VARCHAR(20);not null"`
	Status          PaymentStatus `gorm:"type:VARCHAR(20);default:'PENDING'"`
	TransactionID   string        `gorm:"size:100"` // For digital payments
	Reference       string        `gorm:"size:100"`
	ProcessedBy     *uint         // User who processed
	ProcessedByUser *User         `gorm:"foreignKey:ProcessedBy"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
