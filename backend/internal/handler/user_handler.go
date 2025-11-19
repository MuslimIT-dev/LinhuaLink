package handler

import (
	"LinhuaLink/backend/internal/middleware"
	"LinhuaLink/backend/internal/service"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service service.UserService
}

type UserHandler interface {
	Me(cxt *gin.Context)
}

func NewUserHandler(service service.UserService) UserHandler {
	return &userHandler{service: service}
}

// Me godoc
// @Summary Получение данных текущего пользователя
// @Description Возвращает информацию о пользователе по JWT cookie
// @Tags User
// @Security BearerAuth
// @Produce json
// @Success 200 {object} model.User "Success"
// @Failure 401 {object} model.ErrorResponseCode401 "StatusUnauthorized"
// @Failure 500 {object} model.ErrorResponseCode500 "StatusInternalServerError"
// @Router /user/me [get]
func (h *userHandler) Me(cxt *gin.Context) {
	userId := middleware.GetUserIDFromContext(cxt)
	user, err := h.service.Me(userId)
	if err != nil {
		cxt.JSON(500, gin.H{"error": err.Error()})
		return
	}
	cxt.JSON(200, gin.H{"status": true, "user": user})
}
