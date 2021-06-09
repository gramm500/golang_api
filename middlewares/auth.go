package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("authorization")
		headerSlice := strings.Split(authHeader, " ")
		if len(headerSlice) != 2 {
			c.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}
	}
}
