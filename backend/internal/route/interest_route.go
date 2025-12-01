package route

import (
	"LinhuaLink/backend/internal/handler"
	"LinhuaLink/backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterInterestRoutes(cxt *gin.RouterGroup, h handler.InterestHandler) {
	interest := cxt.Group("/interest", middleware.JWTAuthMiddleware())
	{
		interest.GET("/", h.GetInterest)
	}
}
