package dto

import "time"

// ============================================================================
// USER REQUEST/RESPONSE STRUCTS
// ============================================================================

// CreateUserRequest represents create user request
type CreateUserRequest struct {
	Name         string  `json:"name" validate:"required,max=100"`
	Email        string  `json:"email" validate:"required,email"`
	Password     string  `json:"password" validate:"required,min=6"`
	Phone        string  `json:"phone" validate:"max=20"`
	UserType     string  `json:"user_type" validate:"required,oneof=SUPER_ADMIN RESTAURANT EMPLOYEE"`
	Role         *string `json:"role,omitempty" validate:"omitempty,oneof=WAITER CHEF CASHIER MANAGER KITCHEN_STAFF HOST"`
	RestaurantID *uint   `json:"restaurant_id,omitempty"`
	BranchID     *uint   `json:"branch_id,omitempty"`
	Access       string  `json:"access,omitempty"`
}

// ValidationErrorMessages maps CreateUserRequest fields to custom messages
var CreateUserValidationErrorMessages = map[string]string{
	"Name":     "Name is required and must be at most 100 characters.",
	"Email":    "A valid email is required.",
	"Password": "Password is required and must be at least 6 characters.",
	"UserType": "User type is required and must be one of: SUPER_ADMIN, RESTAURANT, EMPLOYEE.",
	"Role":     "Role must be one of: WAITER, CHEF, CASHIER, MANAGER, KITCHEN_STAFF, HOST.",
	"Phone":    "Phone must be at most 20 characters.",
}

// UpdateUserRequest represents update user request
type UpdateUserRequest struct {
	Name     *string `json:"name,omitempty" validate:"omitempty,max=100"`
	Email    *string `json:"email,omitempty" validate:"omitempty,email"`
	Phone    *string `json:"phone,omitempty" validate:"omitempty,max=20"`
	Role     *string `json:"role,omitempty" validate:"omitempty,oneof=WAITER CHEF CASHIER MANAGER KITCHEN_STAFF HOST"`
	BranchID *uint   `json:"branch_id,omitempty"`
	Access   *string `json:"access,omitempty"`
	IsActive *bool   `json:"is_active,omitempty"`
}

// UserResponse represents user response
type UserResponse struct {
	ID           uint       `json:"id"`
	Name         string     `json:"name"`
	Email        string     `json:"email"`
	Phone        string     `json:"phone"`
	UserType     string     `json:"user_type"`
	Role         *string    `json:"role,omitempty"`
	RestaurantID *uint      `json:"restaurant_id,omitempty"`
	BranchID     *uint      `json:"branch_id,omitempty"`
	Access       string     `json:"access"`
	IsActive     bool       `json:"is_active"`
	LastLogin    *time.Time `json:"last_login,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

// UserSummary represents user summary for auth responses
type UserSummary struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	UserType string  `json:"user_type"`
	Role     *string `json:"role,omitempty"`
}
