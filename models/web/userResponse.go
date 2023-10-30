package  web

type UserResponse struct{
	ID			int	   `json:"id" form:"id"`
	Username	string `json:"username" form:"username"`
	Name		string `json:"name" form:"name"`
	Role		string `json:"role" form:"role"`
	Email		string `json:"email" form:"email"`
}

<<<<<<< Updated upstream
type UserLoginResponse struct {
	ID			int	   `json:"id"`
	Username	string `json:"username"`
	Role		string `json:"role"`
	Token		string `json:"token"`
	Password 	string `json:"password"`
}
=======
type AdminLoginResponse struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Token    string `json:"token"`
}
type UserLoginResponse struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}
>>>>>>> Stashed changes
