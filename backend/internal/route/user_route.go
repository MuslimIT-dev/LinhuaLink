package route

import (
	"LinhuaLink/backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(cxt *gin.RouterGroup, h handler.UserHandler) {
	userApi := cxt.Group("/user")
	{
		userApi.POST("/signup", h.Signup)
	}
}
