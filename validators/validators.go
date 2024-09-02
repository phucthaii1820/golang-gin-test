package validators

import (
	"test/inits"
	"test/models"

	"github.com/go-playground/validator/v10"
)

func UniqueTitle(fl validator.FieldLevel) bool {
	result := inits.DB.First(&models.Card{}, "title = ?", fl.Field().String())

	return result.RowsAffected == 0
}