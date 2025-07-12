package dto

import (
	userdto "restaurant_os/internal/api/user/dto"
	"time"
)

// ============================================================================
// AUTH REQUEST/RESPONSE STRUCTS
// ============================================================================

// LoginRequest represents login request payload
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// LoginResponse represents login response
type LoginResponse struct {
	Token        string              `json:"token"`
	RefreshToken string              `json:"refresh_token"`
	User         userdto.UserSummary `json:"user"`
	ExpiresAt    time.Time           `json:"expires_at"`
}

// RefreshTokenRequest represents refresh token request
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// ChangePasswordRequest represents change password request
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,min=6"`
}
