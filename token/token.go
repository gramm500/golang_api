package token

import (
	"github.com/golang-jwt/jwt"
	"golang_api/models"
	"time"
)

func CreateToken(id interface{}, u models.User) (string, error) {
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	tk := models.Token{
		UserID: id,
		Email:  u.Email,
		Role:   u.Role,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
