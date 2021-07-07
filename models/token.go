package models

import "github.com/golang-jwt/jwt"

type Token struct {
	UserID interface{}
	Email  string
	Role   string
	*jwt.StandardClaims
}
