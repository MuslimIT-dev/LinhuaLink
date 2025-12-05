package handler

import (
	"LinhuaLink/backend/internal/service"

	"github.com/gin-gonic/gin"
)

type InterestHandler interface {
	GetInterest(cxt *gin.Context)
}

type interestHandler struct {
	service service.InterestService
}

func NewInterestHandler(service service.InterestService) InterestHandler {
	return &interestHandler{service: service}
}

// GetInterest godoc
// @Summary Получение всех интересов
// @Description возвращает список всех интересов из базы данных
// @Tags Interest
// @Security BearerAuth
// @Produce json
// @Success 200 {object} model.Interest "Success"
// @Failure 401 {object} model.ErrorResponseCode401 "StatusUnauthorized"
// @Failure 500 {object} model.ErrorResponseCode500 "StatusInternalServerError"
// @Router /interest/ [get]
func (h *interestHandler) GetInterest(cxt *gin.Context) {
	interests, err := h.service.GetInterest()
	if err != nil {
		cxt.JSON(500, gin.H{"error": err.Error()})
		return
	}
	cxt.JSON(200, gin.H{"interests": interests})
}
