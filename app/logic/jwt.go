package logic

import (
	"fmt"
	"github.com/soonio/pupil/app"
	"time"

	"github.com/golang-jwt/jwt"
)

const BlackListCachePrefix = "blacklist:"

type UserClaims struct {
	User struct {
		ID uint `json:"id"`
	} `json:"user" jwt:"user"`
	jwt.StandardClaims
}

// jwt 简单代理
type jwtLogic struct{}

var Jwt = new(jwtLogic)

// Parse 解析(当前并不使用，只是为了演示)
func (t *jwtLogic) Parse(ciphertext string) *UserClaims {
	Token, err := jwt.ParseWithClaims(ciphertext, &UserClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v ", token.Header["alg"])
		}
		return []byte(app.Config.JWT.SigningKey), nil
	})
	if err == nil {
		if claims, ok := Token.Claims.(*UserClaims); ok && Token.Valid {
			return claims
		}
	}
	return nil
}

// Gen 生成jwt类型的token
func (t *jwtLogic) Gen(claims jwt.Claims) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString([]byte(app.Config.JWT.SigningKey))
}

// BlackList 加入黑名单
func (t *jwtLogic) BlackList(UUID string) error {
	return Cache.Set(BlackListCachePrefix+UUID, "", time.Duration(app.Config.JWT.Duration)*time.Second)
}

func (t *jwtLogic) InBlackList(UUID string) (bool, error) {
	return Cache.Has(BlackListCachePrefix + UUID)
}
