package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
<<<<<<< Updated upstream
	Username	string 
	Password	string 
	Name		string 
	Role		string 
	Email		string 
=======
	Username string `gorm:"unique" json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
	Role     Role   `gorm:"type:enum('Admin', 'User')" json:"Role"`
>>>>>>> Stashed changes
}

