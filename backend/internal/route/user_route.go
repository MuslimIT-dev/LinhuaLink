package route

import (
	"LinhuaLink/backend/internal/handler"
	"LinhuaLink/backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(cxt *gin.RouterGroup, h handler.UserHandler) {
	userApi := cxt.Group("/user", middleware.JWTAuthMiddleware())
	{
		userApi.GET("/me", h.Me)
	}
}
