package handlers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"golang_api/models"
	"golang_api/repositories"
	"net/http"
	"time"
)

type RegisterForm struct {
	Email    string `form:"email" binding:"required"`
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

	res, err := h.MongoRepository.CreateUser(&u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "something went horribly wrong")
		return
	}

	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	tk := models.Token{
		UserID: res.InsertedID,
		Email:  u.Email,
		Role:   u.Role,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(http.StatusCreated, tokenString)
}
