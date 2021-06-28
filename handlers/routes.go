package handlers

import (
	"github.com/gin-gonic/gin"
	"golang_api/models"
)

func RegisterRoutes(r *gin.Engine, di *models.DI) {
	registerHandler := NewRegisterHandler(di)
	LoginHandler := NewLoginHandler(di)
	r.POST("/register", registerHandler.Register)
	r.POST("/login", LoginHandler.Login)
}
