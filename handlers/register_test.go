package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"golang_api/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterHandler_Register(t *testing.T) {

	err := godotenv.Load("./../.env")
	if err != nil {
		panic(err)
	}

	var config models.Config
	err = envconfig.Process("", &config)
	if err != nil {
		panic(err)
	}

	testRouter := gin.New()
	testDi := models.NewDI(config)
	RegisterRoutes(testRouter, testDi)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/hi", nil)

	testRouter.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
