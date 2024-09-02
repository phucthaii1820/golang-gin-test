package migrations

import (
	"test/inits"
	"test/models"
)




func Migrate()  {
	inits.DB.AutoMigrate(&models.Card{})
}