package handlers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"golang_api/models"
	"golang_api/repositories"
	"golang_api/token"
	"net/http"
)

type RegisterForm struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type RegisterHandler struct {
	MongoRepository *repositories.MongoDbRepository
}

func NewRegisterHandler(di *models.DI) *RegisterHandler {
	return &RegisterHandler{MongoRepository: di.MongoRepository}
}

func (h *RegisterHandler) Register(ctx *gin.Context) {
	req := &RegisterForm{}
	err := ctx.ShouldBind(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	passwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Something went horribly wrong")
		return
	}

	u := models.User{
		Email:    req.Email,
		Password: string(passwd),
		Role:     "admin",
	}

	res, err := h.MongoRepository.CreateUser(u.Email, u.Password, u.Role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "something went horribly wrong")
		return
	}

	tokenString, err := token.CreateToken(res.InsertedID, u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "something went horribly wrong")
	}

	ctx.JSON(http.StatusCreated, "token: "+tokenString)
}
