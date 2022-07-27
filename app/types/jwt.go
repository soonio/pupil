package types

import "github.com/golang-jwt/jwt"

type UserClaims struct {
	User struct {
		ID uint `json:"id"`
	} `json:"user" jwt:"user"`
	jwt.StandardClaims
}
