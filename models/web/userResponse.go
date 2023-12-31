package web

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username" form:"username"`
	Name     string `json:"name" form:"name"`
	Role     string `json:"role" form:"role"`
}

type AdminLoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	Role     string `json:"role"`
	Name     string `json:"name"`
}

type UserLoginResponse struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Name     string `json:"name"`
}
