package req

import (
		"SaveHouse/models/web"
		"SaveHouse/models"
)

func PassBody(users web.UserRequest) *models.User{
	return &models.User{
		Name	: users.Name,
		Username: users.Username,
		Role	: users.Role,
		Email	: users.Email,
		Password: users.Password,
	}
}