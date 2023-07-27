package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name" form:"name"`
	Description string  `json:"description" form:"description"`
	Price       float64 `json:"price" form:"price" binding:"gte=0"`
	Stock       int     `json:"stock" form:"stock" binding:"gte=0"`
	Rating      float64 `json:"rating" form:"rating" binding:"gte=1,lte=5"`
	Category    string  `json:"category" form:"category"`
}
