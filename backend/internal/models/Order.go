package models

import (
	"gorm.io/gorm"
	"time"
)

type OrderStatus string
type OrderType string
type OrderSource string
type PaymentStatus string

const (
	OrderPending   OrderStatus = "PENDING"
	OrderConfirmed OrderStatus = "CONFIRMED"
	OrderPreparing OrderStatus = "PREPARING"
	OrderReady     OrderStatus = "READY"
	OrderServed    OrderStatus = "SERVED"
	OrderCompleted OrderStatus = "COMPLETED"
	OrderCancelled OrderStatus = "CANCELLED"
	OrderRefunded  OrderStatus = "REFUNDED"
)

const (
	OrderTypeDineIn   OrderType = "DINE_IN"
	OrderTypeTakeaway OrderType = "TAKEAWAY"
	OrderTypeDelivery OrderType = "DELIVERY"
	OrderTypeOnline   OrderType = "ONLINE"
)

const (
	OrderSourceStaff   OrderSource = "STAFF"      // Order taken by waiter/cashier
	OrderSourceQR      OrderSource = "QR_CODE"    // Customer ordered via QR code
	OrderSourceApp     OrderSource = "MOBILE_APP" // Customer ordered via mobile app
	OrderSourceWebsite OrderSource = "WEBSITE"    // Customer ordered via website
	OrderSourceKiosk   OrderSource = "KIOSK"      // Self-service kiosk
)

const (
	PaymentPending  PaymentStatus = "PENDING"
	PaymentPaid     PaymentStatus = "PAID"
	PaymentPartial  PaymentStatus = "PARTIAL"
	PaymentRefunded PaymentStatus = "REFUNDED"
	PaymentFailed   PaymentStatus = "FAILED"
)

type Order struct {
	ID            uint   `gorm:"primaryKey"`
	OrderNumber   string `gorm:"unique;not null;size:50"` // Human-readable order number
	TableID       *uint
	Table         *Table        `gorm:"foreignKey:TableID"`
	BranchID      uint          `gorm:"not null"`
	Branch        Branch        `gorm:"foreignKey:BranchID"`
	UserID        *uint         // Waiter/Cashier who placed (null for QR orders)
	User          *User         `gorm:"foreignKey:UserID"`
	CustomerName  string        `gorm:"size:100"`
	CustomerPhone string        `gorm:"size:20"`
	CustomerEmail string        `gorm:"size:255"`
	OrderType     OrderType     `gorm:"type:VARCHAR(20);default:'DINE_IN'"`
	OrderSource   OrderSource   `gorm:"type:VARCHAR(20);default:'STAFF'"` // How order was placed
	Status        OrderStatus   `gorm:"type:VARCHAR(20);default:'PENDING'"`
	PaymentStatus PaymentStatus `gorm:"type:VARCHAR(20);default:'PENDING'"`

	// QR Order specific fields
	QRSessionID *uint      // Link to QR session if applicable
	QRSession   *QRSession `gorm:"foreignKey:QRSessionID"`
	IsQROrder   bool       `gorm:"default:false"`

	// Financial fields
	Subtotal       float64 `gorm:"type:decimal(10,2);default:0"`
	TaxAmount      float64 `gorm:"type:decimal(10,2);default:0"`
	DiscountAmount float64 `gorm:"type:decimal(10,2);default:0"`
	ServiceCharge  float64 `gorm:"type:decimal(10,2);default:0"`
	Total          float64 `gorm:"type:decimal(10,2);default:0"`

	Notes         string `gorm:"type:text"`
	EstimatedTime int    `gorm:"default:0"` // minutes
	OrderItems    []OrderItem
	Payments      []Payment

	// Staff assignment for QR orders
	AssignedWaiterID *uint // Waiter assigned to serve this QR order
	AssignedWaiter   *User `gorm:"foreignKey:AssignedWaiterID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
