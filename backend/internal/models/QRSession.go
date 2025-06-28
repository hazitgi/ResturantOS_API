package models

import (
	"gorm.io/gorm"
	"time"
)

type QRSessionStatus string

const (
	QRSessionActive    QRSessionStatus = "ACTIVE"
	QRSessionCompleted QRSessionStatus = "COMPLETED"
	QRSessionExpired   QRSessionStatus = "EXPIRED"
	QRSessionAbandoned QRSessionStatus = "ABANDONED"
)

type QRSession struct {
	ID            uint            `gorm:"primaryKey"`
	SessionToken  string          `gorm:"unique;not null;size:255"` // Unique session token
	TableID       uint            `gorm:"not null"`
	Table         Table           `gorm:"foreignKey:TableID"`
	BranchID      uint            `gorm:"not null"`
	Branch        Branch          `gorm:"foreignKey:BranchID"`
	CustomerName  string          `gorm:"size:100"`
	CustomerPhone string          `gorm:"size:20"`
	CustomerEmail string          `gorm:"size:255"`
	GuestCount    int             `gorm:"default:1"`
	Status        QRSessionStatus `gorm:"type:VARCHAR(20);default:'ACTIVE'"`

	// Session management
	StartedAt      time.Time `gorm:"not null"`
	LastActivityAt time.Time `gorm:"not null"`
	ExpiresAt      time.Time `gorm:"not null"` // Session expiry time

	// Cart functionality for QR orders
	CartItems []QRCartItem // Items in cart before order placement
	Orders    []Order      // Orders placed in this session

	// Device information
	DeviceInfo string `gorm:"type:text"` // JSON with device/browser info
	IPAddress  string `gorm:"size:45"`   // IPv4 or IPv6

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// QR Cart Item model (new) - for cart functionality in QR ordering
type QRCartItem struct {
	ID          uint      `gorm:"primaryKey"`
	QRSessionID uint      `gorm:"not null"`
	QRSession   QRSession `gorm:"foreignKey:QRSessionID"`
	MenuItemID  uint      `gorm:"not null"`
	MenuItem    MenuItem  `gorm:"foreignKey:MenuItemID"`
	Quantity    int       `gorm:"not null;default:1"`
	UnitPrice   float64   `gorm:"type:decimal(10,2);not null"`
	TotalPrice  float64   `gorm:"type:decimal(10,2);not null"`
	Notes       string    `gorm:"type:text"` // Special instructions
	AddedAt     time.Time `gorm:"not null"`
	UpdatedAt   time.Time
}

// QR Code Analytics model (new) - for tracking QR code usage
type QRCodeScan struct {
	ID          uint       `gorm:"primaryKey"`
	TableID     uint       `gorm:"not null"`
	Table       Table      `gorm:"foreignKey:TableID"`
	BranchID    uint       `gorm:"not null"`
	Branch      Branch     `gorm:"foreignKey:BranchID"`
	QRSessionID *uint      // If scan led to a session
	QRSession   *QRSession `gorm:"foreignKey:QRSessionID"`

	// Analytics data
	IPAddress  string    `gorm:"size:45"`
	UserAgent  string    `gorm:"type:text"`
	DeviceType string    `gorm:"size:50"` // mobile, tablet, desktop
	ScanTime   time.Time `gorm:"not null"`

	// Conversion tracking
	ConvertedToSession bool `gorm:"default:false"`
	ConvertedToOrder   bool `gorm:"default:false"`

	CreatedAt time.Time
}
