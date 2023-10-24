package  web

type UserResponse struct{
	ID			int	   `json:"id" form:"id"`
	Username	string `json:"username" form:"username"`
	Name		string `json:"name" form:"name"`
	Role		string `json:"role" form:"role"`
	Email		string `json:"email" form:"email"`
}

type UserLoginResponse struct {
	ID			int	   `json:"id"`
	Username	string `json:"username"`
	Role		string `json:"role"`
	Token		string `json:"token"`
	Password 	string `json:"password"`
}