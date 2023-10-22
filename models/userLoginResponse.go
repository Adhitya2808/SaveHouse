package models

type UserLoginResponse struct {
	ID			int	   `gorm:"primaryKey;type:smallint" json:"id" form:"id"`
	Username	string `gorm:"type:varchar(255)" json:"username" form:"username"`
	Role		string `gorm:"type:varchar(5)" json:"role" form:"role"`
	Token		string `gorm:"-" json:"token" form:"token"`
}
