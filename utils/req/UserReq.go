package req

import (
	"SaveHouse/models"
	"SaveHouse/models/web"
)

func PassBody(users web.UserRequest) *models.User {
	return &models.User{
		Name:     users.Name,
		Username: users.Username,
		Password: users.Password,
		Role:     users.Role,
	}
}
