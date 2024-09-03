package middlewares

import (
	"strings"
	"test/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		token := strings.Split(c.GetHeader("Authorization"), "Bearer ")[1]
		
		if token == "" {
			c.JSON(401, gin.H{"error": "Authorization token not provided"})
			c.Abort()

			return 
		}

		if !utils.ValidateToken(token) {
            c.JSON(401, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }
		
		c.Next()
	}
}