package main

import (
	"LinhuaLink/backend/internal/handler"
	"LinhuaLink/backend/internal/repository"
	"LinhuaLink/backend/internal/route"
	"LinhuaLink/backend/internal/service"
	"LinhuaLink/backend/pkg/config"
	"LinhuaLink/backend/pkg/database"

	_ "LinhuaLink/backend/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title        LinhuaLink API
// @version      1.0
// @description  API документация для проекта LinhuaLink.
// @host         localhost:8080
// @BasePath     /api

func main() {
	config.LoadEnv()
	database.InitDB()

	c := gin.Default()

	authRepository := repository.NewAuthRepository(database.DB)
	authService := service.NewAuthService(authRepository)
	authHandler := handler.NewAuthHandler(authService)

	userRepository := repository.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	interestRepository := repository.NewInterestRepository(database.DB)
	interestService := service.NewInterestService(interestRepository)
	interestHandler := handler.NewInterestHandler(interestService)

	api := c.Group("/api")
	{
		route.RegisterAuthRoutes(api, authHandler)
		route.RegisterUserRoutes(api, userHandler)
		route.RegisterInterestRoutes(api, interestHandler)
	}

	c.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	c.Run(":8080")
}
