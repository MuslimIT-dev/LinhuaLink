package middleware

import (
	"LinhuaLink/backend/pkg/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const contextUserId = "userId"

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := c.Cookie("token")
		fmt.Println(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token"})
			c.Abort()
			return
		}

		token, err := utils.ParseAndVerifyJWT(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userId", int(token["userId"].(float64)))
		c.Next()
	}
}

func GetUserIDFromContext(c *gin.Context) int {
	id, exists := c.Get(contextUserId)
	if !exists {
		return 0
	}

	return id.(int)
}
