package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"golang_api/controllers"
	"golang_api/models"
	"log"
)

func main() {
	err := godotenv.Load("./../.env")
	if err != nil {
		panic(err)
	}

	var config models.Config
	err = envconfig.Process("", &config)
	if err != nil {
		panic(err)
	}

	di := models.NewDI(config)

	router := SetupRouter(di)

	router.POST("/register", controllers.Register)
	router.GET("/hi", controllers.Welcome)
	log.Fatal(router.Run(":3000"))
}

func SetupRouter(di *models.DI) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	handlers.RegisterRoutes(r, di)
	r.NoRoute(defaultHandler.Default)
	return r
}
