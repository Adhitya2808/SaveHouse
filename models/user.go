package models

type User struct {
	ID			int	   `gorm:"primary_key;type:smallint" json:"id" form:"id"`
	Username	string `gorm:"type:varchar(255);not null" json:"username" form:"username"`
	Password	string `gorm:"type:varchar(225);not null" json:"password" form:"password"`
	Name		string `gorm:"type:varchar(255);not null" json:"name" form:"name"`
	Role		string `gorm:"type:varchar(5);not null" json:"role" form:"role"`
	Email		string `gorm:"type:varchar(255);not null" json:"email" form:"email"`
}

