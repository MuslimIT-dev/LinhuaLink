package route

import (
	"LinhuaLink/backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(cxt *gin.RouterGroup, h handler.AuthHandler) {
	userApi := cxt.Group("/auth")
	{
		userApi.GET("/login", h.Login)
		userApi.POST("/signup", h.Signup)
		userApi.POST("/logout", h.Logout)
	}
}
