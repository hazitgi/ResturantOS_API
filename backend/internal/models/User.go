package models

import (
	"gorm.io/gorm"
	"time"
)

// User Types
type UserType string

const (
	UserTypeSuperAdmin UserType = "SUPER_ADMIN"
	UserTypeRestaurant UserType = "RESTAURANT"
	UserTypeEmployee   UserType = "EMPLOYEE"
)

// Employee Roles
type EmployeeRole string

const (
	RoleWaiter  EmployeeRole = "WAITER"
	RoleChef    EmployeeRole = "CHEF"
	RoleCashier EmployeeRole = "CASHIER"
	RoleManager EmployeeRole = "MANAGER"
	RoleKitchen EmployeeRole = "KITCHEN_STAFF"
	RoleHost    EmployeeRole = "HOST"
)

// User model with enhanced fields
type User struct {
	ID           uint          `gorm:"primaryKey"`
	Name         string        `gorm:"not null;size:100"`
	Email        string        `gorm:"unique;not null;size:255"`
	Password     string        `gorm:"not null;size:255"`
	Phone        string        `gorm:"size:20"`
	UserType     UserType      `gorm:"type:VARCHAR(20);not null"`
	Role         *EmployeeRole `gorm:"type:VARCHAR(20)"`
	RestaurantID *uint
	Restaurant   *Restaurant `gorm:"foreignKey:RestaurantID"`
	BranchID     *uint
	Branch       *Branch `gorm:"foreignKey:BranchID"`
	Access       string  `gorm:"type:text"`
	IsActive     bool    `gorm:"default:true"`
	LastLogin    *time.Time
	CreatedBy    *uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
