package models

import jwt "github.com/dgrijalva/jwt-go"

type Token struct {
	UserID interface{}
	Email  string
	Role   string
	*jwt.StandardClaims
}
