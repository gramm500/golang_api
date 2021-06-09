package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterRequest struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Register(ctx *gin.Context) {
	req := &RegisterRequest{}
	err := ctx.ShouldBindQuery(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	ctx.JSON(http.StatusOK, "all good homie")
}
