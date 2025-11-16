package handler

import (
	"LinhuaLink/backend/internal/model"
	"LinhuaLink/backend/internal/service"
	"LinhuaLink/backend/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service service.UserService
}

type UserHandler interface {
	Signup(cxt *gin.Context)
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
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		cxt.JSON(500, gin.H{"error": err.Error()})
	}
	cxt.SetCookie("token", token, int((time.Hour * 24).Seconds()), "/", "", false, true) // 24 hours

	cxt.JSON(200, gin.H{"status": true, "user": req})
}
