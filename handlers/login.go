package handlers

import (
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	//var u models.User
	//if err := c.ShouldBindJSON(&u); err != nil {
	//	c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
	//	return
	//}
	////compare the user from the request, with the one we defined:
	//if user.Username != u.Username || user.Password != u.Password {
	//	c.JSON(http.StatusUnauthorized, "Please provide valid login details")
	//	return
	//}
	//token, err := CreateToken(user.ID)
	//if err != nil {
	//	c.JSON(http.StatusUnprocessableEntity, err.Error())
	//	return
	//}
	//c.JSON(http.StatusOK, token)
}