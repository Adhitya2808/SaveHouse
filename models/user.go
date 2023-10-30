package models

import "gorm.io/gorm"

type Role string

const (
	AdminRole = "Admin"
	UserRole  = "User"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
	Role     Role   `gorm:"type:enum('Admin', 'User')" json:"Role"`
}
