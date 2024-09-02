package routers

import (
	"test/controllers"

	"github.com/gin-gonic/gin"
)



func CardRouter(r *gin.Engine) {

	r.GET("/card/:id",  controllers.GetCardById)
	r.POST("/card", controllers.CreateCard)
	r.DELETE("/card/:id", controllers.DeleteCard)
	r.GET("/cards", controllers.GetCards)
	r.PUT("/card/:id", controllers.UpdateCard)

}