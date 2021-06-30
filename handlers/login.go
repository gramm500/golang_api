package handlers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"golang_api/models"
	"golang_api/repositories"
	"golang_api/token"
	"net/http"
)

type LoginRequest struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type LoginHandler struct {
	MongoRepository *repositories.MongoDbRepository
}

func NewLoginHandler(di *models.DI) *LoginHandler {
	return &LoginHandler{MongoRepository: di.MongoRepository}
}

func (h *LoginHandler) Login(ctx *gin.Context) {

	req := &RegisterForm{}
	err := ctx.ShouldBind(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	var user models.User
	err = h.MongoRepository.FindUser(req.Email).Decode(&user)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "no users with this email found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "wrong password")
	}

	userToken, err := token.CreateToken(user.Id, user)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, userToken)
}
