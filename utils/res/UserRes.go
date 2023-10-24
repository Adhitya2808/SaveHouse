package res

import (
	"SaveHouse/models"
	"SaveHouse/models/web"
)

func ConvertIndex(users []models.User) []web.UserResponse {
	var results []web.UserResponse
	for _, user := range users {
		userResponse := web.UserResponse{
			ID:       int(user.ID),
			Name:     user.Name,
			Username: user.Username,
			Email:    user.Email,
			Role: 	  user.Role,	
		}
		results = append(results, userResponse)
	}

	return results
}

func ConvertGeneral(user *models.User) web.UserResponse {
	return web.UserResponse{
		ID:       int(user.ID),
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Role: 	  user.Role,	
	}
}