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

func (h *interestHandler) GetInterest(cxt *gin.Context) {
	interests, err := h.service.GetInterest()
	if err != nil {
		cxt.JSON(500, gin.H{"error": err.Error()})
		return
	}
	cxt.JSON(200, gin.H{"interests": interests})
}
