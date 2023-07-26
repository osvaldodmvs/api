package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" binding:"required,email" gorm:"unique"`
	Password string `json:"password,omitempty" binding:"required,min=6" `
}
