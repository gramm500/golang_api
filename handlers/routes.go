package handlers

import (
	"github.com/gin-gonic/gin"
	"golang_api/models"
)

func RegisterRoutes(r *gin.Engine, di *models.DI) {
	registerHandler := NewRegisterHandler(di)
	r.POST("/register", registerHandler.Register)
	r.GET("/hi", Welcome)
}
