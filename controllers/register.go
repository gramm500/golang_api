package controllers

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"golang_api/models"
	"net/http"
	"time"
)

type RegisterForm struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Register(ctx *gin.Context) {
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
	fmt.Println(passwd)
	u := models.User{
		Email:    req.Email,
		Password: string(passwd),
		Role:     "admin",
	}

	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Something went horribly wrong")
		return
	}
	collection := client.Database("go_api").Collection("users")
	res, err := collection.InsertOne(ctx, bson.D{{"email", u.Email}, {"password", u.Password}, {"role", u.Role}})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Something went horribly wrong")
		return
	}
	id := res.InsertedID
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	tk := models.Token{
		UserID: id,
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
