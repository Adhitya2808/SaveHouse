package res

import (
	"app/models"
	"app/models/web"
)

func ConvertIndex(users []models.User) []web.UserResponse {
	var results []web.UserResponse
	for _, user := range users {
		userResponse := web.UserResponse{
			ID:       int(user.ID),
			Name:     user.Name,
			Username: user.Username,
			Role:     string(user.Role),
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
		Role:     string(user.Role),
	}
}
