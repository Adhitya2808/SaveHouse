package models

type UserResponse struct{
	ID			int	   `gorm:"primaryKey;type:smallint not null" json:"id" form:"id"`
	Username	string `gorm:"type:varchar(255) not null" json:"username" form:"username"`
	Name		string `gorm:"type:varchar(255) not null" json:"name" form:"name"`
	Role		string `gorm:"type:varchar(5) not null" json:"role" form:"role"`
	Email		string `gorm:"type:varchar(255) not null" json:"email" form:"email"`
}