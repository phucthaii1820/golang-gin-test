package models

import (
	"gorm.io/gorm"
)



type Card struct {
	gorm.Model
	
	Title string
	Author string
}

type CreateCardInput struct {
	Title string `json:"title" binding:"required" validate:"uniqueTitle"`
	Author string `json:"author" binding:"required"`
}
