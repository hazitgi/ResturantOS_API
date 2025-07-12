package services

import (
	"fmt"
	"restaurant_os/internal/api/user/dto"
	"restaurant_os/internal/models"
)

func CreateUser(userDetails *dto.CreateUserRequest) (*models.User, error) {
	// check is user exists
	existingUser := &models.User{}
	if err := models.DataBase.Where("email = ?", userDetails.Email).First(existingUser).Error; err == nil {
		return nil, fmt.Errorf("user already exists") // Assuming this error is defined in models
	}

	user := &models.User{
		Name:     userDetails.Name,
		Email:    userDetails.Email,
		Password: userDetails.Password, // Hash before saving in real use
		Phone:    userDetails.Phone,
		UserType: models.UserType(userDetails.UserType),
		Access:   userDetails.Access,
		IsActive: true,
	}
	if userDetails.Role != nil {
		role := models.EmployeeRole(*userDetails.Role)
		user.Role = &role
	}
	if userDetails.RestaurantID != nil {
		user.RestaurantID = userDetails.RestaurantID
	}
	if userDetails.BranchID != nil {
		user.BranchID = userDetails.BranchID
	}

	// Save user to DB (assumes models.DataBase is initialized)
	if err := models.DataBase.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
