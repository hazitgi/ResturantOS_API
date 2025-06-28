package dto

import "github.com/golang-jwt/jwt/v5"

// ============================================================================
// COMMON STRUCTS
// ============================================================================

// APIResponse represents a standard API response structure
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   *string     `json:"error,omitempty"`
}

// PaginatedResponse represents a paginated API response
type PaginatedResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination *Pagination `json:"pagination"`
	Error      *string     `json:"error,omitempty"`
}

// Pagination represents pagination metadata
type Pagination struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
	HasNext    bool  `json:"has_next"`
	HasPrev    bool  `json:"has_prev"`
}

// Claims represents JWT claims
type Claims struct {
	UserID       uint   `json:"user_id"`
	Email        string `json:"email"`
	UserType     string `json:"user_type"`
	Role         string `json:"role,omitempty"`
	RestaurantID *uint  `json:"restaurant_id,omitempty"`
	BranchID     *uint  `json:"branch_id,omitempty"`
	jwt.RegisteredClaims
}
