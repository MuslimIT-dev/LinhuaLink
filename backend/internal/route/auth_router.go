package route

import (
	"LinhuaLink/backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(cxt *gin.RouterGroup, h handler.AuthHandler) {
	userApi := cxt.Group("/auth")
	{
		userApi.POST("/signup", h.Signup)
		userApi.GET("/login", h.Login)
	}
}
