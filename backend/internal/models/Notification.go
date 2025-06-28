package models

import (
	"time"
)

type NotificationType string
type NotificationStatus string

const (
	NotificationNewOrder       NotificationType = "NEW_ORDER"
	NotificationOrderReady     NotificationType = "ORDER_READY"
	NotificationPaymentPending NotificationType = "PAYMENT_PENDING"
	NotificationQRSession      NotificationType = "QR_SESSION"
	NotificationInventoryLow   NotificationType = "INVENTORY_LOW"
	NotificationReservation    NotificationType = "RESERVATION"
)

const (
	NotificationPending NotificationStatus = "PENDING"
	NotificationSent    NotificationStatus = "SENT"
	NotificationRead    NotificationStatus = "READ"
	NotificationFailed  NotificationStatus = "FAILED"
)

type Notification struct {
	ID       uint               `gorm:"primaryKey"`
	BranchID uint               `gorm:"not null"`
	Branch   Branch             `gorm:"foreignKey:BranchID"`
	UserID   *uint              // Target user (null for broadcast)
	User     *User              `gorm:"foreignKey:UserID"`
	Type     NotificationType   `gorm:"type:VARCHAR(30);not null"`
	Status   NotificationStatus `gorm:"type:VARCHAR(20);default:'PENDING'"`
	Title    string             `gorm:"not null;size:200"`
	Message  string             `gorm:"type:text;not null"`
	Data     string             `gorm:"type:json"` // Additional data as JSON

	// Reference IDs for related entities
	OrderID     *uint
	TableID     *uint
	QRSessionID *uint

	ReadAt    *time.Time
	SentAt    *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
