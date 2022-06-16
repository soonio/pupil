package tools

import (
	"fmt"
	"time"

	"github.com/soonio/pupil/app/logic"

	"github.com/golang-jwt/jwt"
	"github.com/urfave/cli/v2"
)

type frontUser struct {
	ID uint `json:"id"`
}

func GenJwt(c *cli.Context) error {

	id := c.Uint("id")

	token, _ := logic.Jwt.Gen(&logic.UserClaims{
		User:           frontUser{ID: id},
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Unix() + 8640000, Issuer: "api"},
	})
	fmt.Printf("uid[%d] front[%s] \n", id, token)

	token, _ = logic.Jwt.Gen(&logic.UserClaims{
		User:           frontUser{ID: id},
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Unix() + 8640000, Issuer: "admin"},
	})
	fmt.Printf("uid[%d] admin[%s] \n", id, token)
	return nil
}
