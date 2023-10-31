package web

import "app/models"

type UserRequest struct {
	Name     string      `json:"name" form:"name"`
	Username string      `json:"username" form:"username"`
	Role     models.Role `json:"role" form:"role"`
	Password string      `json:"password" form:"password"`
}

type UserLoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
