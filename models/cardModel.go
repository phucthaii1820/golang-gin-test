package models

import "gorm.io/gorm"


type Card struct {
	gorm.Model
	
	Title string
	Author string
}

type CreateCardInput struct {
	Title string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}
