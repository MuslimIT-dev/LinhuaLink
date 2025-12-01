package handler

import (
	"LinhuaLink/backend/internal/model"
	"LinhuaLink/backend/internal/service"
	"LinhuaLink/backend/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	service service.AuthService
}

type AuthHandler interface {
	Signup(cxt *gin.Context)
	Login(cxt *gin.Context)
	Logout(cxt *gin.Context)
}

func NewAuthHandler(service service.AuthService) AuthHandler {
	return &authHandler{service: service}
}

// Signup godoc
// @Summary Регистрация пользователя
// @Description Создаёт нового пользователя и выдаёт JWT cookie
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body model.RegisterUserInput true "Данные для регистрации"
// @Success 200 {object} model.User "Success"
// @Failure 400 {object} model.ErrorResponseCode400 "StatusBadRequest"
// @Failure 500 {object} model.ErrorResponseCode500 "StatusInternalServerError"
// @Router /auth/signup [post]
func (h *authHandler) Signup(cxt *gin.Context) {
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

// Login godoc
// @Summary Авторизация пользователя
// @Description Логин по email/паролю и выдача JWT cookie
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body model.LoginUserInput true "Данные для логина"
// @Success 200 {object} model.User "Success"
// @Failure 400 {object} model.ErrorResponseCode400 "StatusBadRequest"
// @Failure 500 {object} model.ErrorResponseCode500 "StatusInternalServerError"
// @Router /auth/login [post]
func (h *authHandler) Login(cxt *gin.Context) {
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

// Logout godoc
// @Summary Выход пользователя
// @Description очищает JWT cookie
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} model.User "Success"
// @Router /auth/logout [post]
func (h *authHandler) Logout(cxt *gin.Context) {
	cxt.SetCookie("token", "", -1, "/", "", false, true)
	cxt.JSON(200, gin.H{"status": true, "message": "Logged out successfully"})
}
