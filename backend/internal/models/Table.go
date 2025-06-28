package models

import (
	"gorm.io/gorm"
	"time"
)

type TableStatus string

const (
	TableAvailable TableStatus = "AVAILABLE"
	TableOccupied  TableStatus = "OCCUPIED"
	TableReserved  TableStatus = "RESERVED"
	TableCleaning  TableStatus = "CLEANING"
	TableBlocked   TableStatus = "BLOCKED"
)

type Table struct {
	ID           uint        `gorm:"primaryKey"`
	Number       string      `gorm:"not null;size:20"`
	BranchID     uint        `gorm:"not null"`
	Branch       Branch      `gorm:"foreignKey:BranchID"`
	Capacity     int         `gorm:"not null;default:4"`
	Status       TableStatus `gorm:"type:VARCHAR(20);default:'AVAILABLE'"`
	Location     string      `gorm:"size:100"`        // e.g., "Window side", "Corner"
	QRCode       string      `gorm:"unique;size:255"` // Unique QR code for table
	QRToken      string      `gorm:"unique;size:255"` // Token for QR code authentication
	QRMenuURL    string      `gorm:"size:500"`        // Direct URL to QR menu for this table
	IsQRActive   bool        `gorm:"default:true"`    // Enable/disable QR ordering for table
	Orders       []Order
	Reservations []Reservation
	QRSessions   []QRSession // Active QR ordering sessions
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
