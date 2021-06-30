package handlers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"golang_api/models"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestLoginHandler_Login(t *testing.T) {
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
	buffer := new(bytes.Buffer)
	params := url.Values{}
	params.Set("email", "example@gmail.com")
	params.Set("password", "secret123")
	buffer.WriteString(params.Encode())

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/login", buffer)
	if err != nil {
		panic(err)
	}
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	testRouter.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
