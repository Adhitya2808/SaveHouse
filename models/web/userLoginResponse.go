package web


type UserRequest struct {
	Name 	 string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	Role	 string `json:"role" form:"role"`
	Email	 string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserLoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
