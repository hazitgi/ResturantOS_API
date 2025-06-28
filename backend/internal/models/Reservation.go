package models

import (
	"gorm.io/gorm"
	"time"
)

type ReservationStatus string

const (
	ReservationPending   ReservationStatus = "PENDING"
	ReservationConfirmed ReservationStatus = "CONFIRMED"
	ReservationSeated    ReservationStatus = "SEATED"
	ReservationCompleted ReservationStatus = "COMPLETED"
	ReservationCancelled ReservationStatus = "CANCELLED"
	ReservationNoShow    ReservationStatus = "NO_SHOW"
)

type Reservation struct {
	ID            uint   `gorm:"primaryKey"`
	BranchID      uint   `gorm:"not null"`
	Branch        Branch `gorm:"foreignKey:BranchID"`
	TableID       *uint
	Table         *Table            `gorm:"foreignKey:TableID"`
	CustomerName  string            `gorm:"not null;size:100"`
	CustomerPhone string            `gorm:"not null;size:20"`
	CustomerEmail string            `gorm:"size:255"`
	GuestCount    int               `gorm:"not null;default:1"`
	ReservedDate  time.Time         `gorm:"not null"`
	ReservedTime  time.Time         `gorm:"not null"`
	Duration      int               `gorm:"default:120"` // minutes
	Status        ReservationStatus `gorm:"type:VARCHAR(20);default:'PENDING'"`
	Notes         string            `gorm:"type:text"`
	CreatedBy     *uint
	CreatedByUser *User `gorm:"foreignKey:CreatedBy"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
