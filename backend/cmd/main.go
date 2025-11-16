package main

import (
	"LinhuaLink/backend/internal/handler"
	"LinhuaLink/backend/internal/repository"
	"LinhuaLink/backend/internal/route"
	"LinhuaLink/backend/internal/service"
	"LinhuaLink/backend/pkg/config"
	"LinhuaLink/backend/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.InitDB()

	c := gin.Default()

	userRepository := repository.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	api := c.Group("/")
	{
		route.RegisterUserRoutes(api, userHandler)
	}
	c.Run(":8080")
}
