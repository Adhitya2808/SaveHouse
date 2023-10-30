package models

import "gorm.io/gorm"

type Role string

const (
	AdminRole = "Admin"
	UserRole  = "User"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Name     string
	Role     Role `gorm:"type:enum('Admin', 'User')" json:"Role"`
}
