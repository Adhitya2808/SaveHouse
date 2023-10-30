package web

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username" form:"username"`
	Name     string `json:"name" form:"name"`
	Role     string `json:"role" form:"role"`
}

type UserLoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
