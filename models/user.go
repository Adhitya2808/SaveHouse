package models

import "gorm.io/gorm"

type Role string

const (
	AdminRole = "Admin"
	UserRole  = "User"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"Username" form:"Username"`
	Password string `json:"Password" form:"Password"`
	Name     string `json:"Name" form:"Name"`
	Role     Role   `gorm:"type:enum('Admin', 'User')" json:"Role"`
}
