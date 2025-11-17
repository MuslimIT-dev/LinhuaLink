package handler

import (
	"LinhuaLink/backend/internal/middleware"
	"LinhuaLink/backend/internal/model"
	"LinhuaLink/backend/internal/service"
	"LinhuaLink/backend/pkg/utils"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service service.UserService
}

type UserHandler interface {
	Signup(cxt *gin.Context)
	Login(cxt *gin.Context)
	Me(cxt *gin.Context)
}

func NewUserHandler(service service.UserService) UserHandler {
	return &userHandler{service: service}
}

func (h *userHandler) Signup(cxt *gin.Context) {
	var req model.RegisterUserInput
	if err := cxt.ShouldBindJSON(&req); err != nil {
		cxt.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, err := h.service.SignUp(req)
	if err != nil {
		cxt.JSON(500, gin.H{"error": err.Error()})
		return
	}
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		cxt.JSON(500, gin.H{"error": err.Error()})
	}
	cxt.SetCookie("token", token, int((time.Hour * 24).Seconds()), "/", "", false, true) // 24 hours

	cxt.JSON(200, gin.H{"status": true, "user": req})
}

func (h *userHandler) Login(cxt *gin.Context) {
	var req model.LoginUserInput
	if err := cxt.ShouldBindJSON(&req); err != nil {
		cxt.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, err := h.service.Login(req)
	if err != nil {
		cxt.JSON(500, gin.H{"error": err.Error()})
		return
	}
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		cxt.JSON(500, gin.H{"error": err.Error()})
	}
	cxt.SetCookie("token", token, int((time.Hour * 24).Seconds()), "/", "", false, true)

	cxt.JSON(200, gin.H{"status": true, "user": user})
}

func (h *userHandler) Me(cxt *gin.Context) {
	userId := middleware.GetUserIDFromContext(cxt)
	fmt.Println(userId)
	user, err := h.service.Me(userId)
	if err != nil {
		cxt.JSON(500, gin.H{"error": err.Error()})
		return
	}
	cxt.JSON(200, gin.H{"status": true, "user": user})
}
