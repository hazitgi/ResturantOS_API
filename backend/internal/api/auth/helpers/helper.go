package helpers

import (
	"fmt"
	"restaurant_os/internal/models"
	"time"

	"restaurant_os/internal/config"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateAccessToken(user *models.User) string {
	secret := config.EnvConfig.JWTAccessSecret
	expiryStr := config.EnvConfig.JWTAccessExpiry

	duration := time.Hour
	if d, err := parseDuration(expiryStr); err == nil {
		duration = d
	}
	fmt.Println("user.Role", user)
	claims := jwt.MapClaims{
		"user_id":       user.ID,
		"email":         user.Email,
		"user_type":     user.UserType,
		"role":          user.Role,
		"restaurant_id": user.RestaurantID,
		"branch_id":     user.BranchID,
		"exp":           time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return ""
	}
	return signedToken
}

func GenerateRefreshToken(user *models.User) string {
	secret := config.EnvConfig.JWTRefreshSecret
	expiryStr := config.EnvConfig.JWTRefreshExpiry

	duration := 7 * 24 * time.Hour
	if d, err := parseDuration(expiryStr); err == nil {
		duration = d
	}
	claims := jwt.MapClaims{
		"user_id":       user.ID,
		"email":         user.Email,
		"user_type":     user.UserType,
		"role":          user.Role,
		"restaurant_id": user.RestaurantID,
		"branch_id":     user.BranchID,
		"exp":           time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return ""
	}
	return signedToken
}

// parseDuration parses a string like "1h", "7d" into time.Duration
func parseDuration(s string) (time.Duration, error) {
	if len(s) > 1 && s[len(s)-1] == 'd' {
		days, err := time.ParseDuration(s[:len(s)-1] + "h")
		if err != nil {
			return 0, err
		}
		return days * 24, nil
	}
	return time.ParseDuration(s)
}
