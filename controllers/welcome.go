package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Welcome(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "all good homie")
}
