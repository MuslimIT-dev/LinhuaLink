package route

import (
	"LinhuaLink/backend/internal/handler"
	"LinhuaLink/backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(cxt *gin.RouterGroup, h handler.UserHandler) {
	userApi := cxt.Group("/user")
	{
		userApi.POST("/signup", h.Signup)
		userApi.GET("/login", h.Login)
		userApi.GET("/me", middleware.JWTAuthMiddleware(), h.Me)
	}
}
