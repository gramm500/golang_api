package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang_api/controllers"
	"log"
)

type User struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email string             `json:"email,omitempty" bson:"email,omitempty"`
	Role  string             `json:"role,omitempty" bson:"role,omitempty"`
}

var client *mongo.Client

var (
	router = gin.Default()
)

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	router.POST("/register", controllers.Register)
	router.GET("/hi", controllers.Welcome)
	log.Fatal(router.Run(":3000"))
}
