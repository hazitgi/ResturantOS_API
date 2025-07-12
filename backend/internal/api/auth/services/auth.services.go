package services

import (
	"errors"
	"fmt"
	"restaurant_os/internal/models"

	"gorm.io/gorm"
)

func GetUserByEmail(email string) (*models.User, error) {
	fmt.Println("Fetching user by email:", email)
	var user models.User

	err := models.DataBase.Where("email = ?", email).First(&user).Error
	fmt.Println("User fetched:", user, "Error:", err)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("error fetching user: %w", err)
	}
	return &user, nil

}
