package controllers

import (
	"net/http"
	"test/inits"
	"test/models"
	"test/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


var validate *validator.Validate



func GetCardById(c *gin.Context,) {
	id := c.Param("id")
	var card models.Card
	result := inits.DB.First(&card, id)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error})
	} else {
		c.JSON(200, gin.H{"data": card})
	}

}

func GetCards(c *gin.Context,) {
	var cards []models.Card
	result := inits.DB.Find(&cards)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error})
	} else {
		c.JSON(200, gin.H{"data": cards})
	}
}

func CreateCard(c *gin.Context,) {
	var req models.CreateCardInput

	validate = validator.New()
	validate.RegisterValidation("uniqueTitle", validators.UniqueTitle)
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	card := models.Card{Title: req.Title, Author: req.Author}
	result := inits.DB.Create(&card)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error})
	} else {
		c.JSON(200, gin.H{"data": card})
	}
}

func UpdateCard(c *gin.Context,) {
	id := c.Param("id")
	var card models.Card
	result := inits.DB.First(&card, id)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error})
		return
	}

	var req models.CreateCardInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	card.Title = req.Title
	card.Author = req.Author
	result = inits.DB.Save(&card)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error})
	} else {
		c.JSON(200, gin.H{"data": card})
	}
}


func DeleteCard(c *gin.Context,) {
	id := c.Param("id")
	var card models.Card
	result := inits.DB.Delete(&card, id)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error})
	} else {
		c.JSON(200, gin.H{"data": card})
	}
	
}