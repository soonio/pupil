package route

import (
	"github.com/soonio/pupil/app"
	"github.com/soonio/pupil/app/types"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// User 解析jwt密钥
func configure() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:      &types.UserClaims{},
		SigningKey:  []byte(app.Config.JWT.SigningKey),
		TokenLookup: "header:token",
	})
}

func user() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			claims := token.Claims.(*types.UserClaims)
			if claims.User.ID > 0 {
				return next(c)
			}
			return &echo.HTTPError{
				Code:    middleware.ErrJWTInvalid.Code,
				Message: middleware.ErrJWTInvalid.Message,
			}
		}
	}
}

func administrator() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			claims := token.Claims.(*types.UserClaims)
			if claims.User.ID > 0 && claims.Issuer == "admin" {
				c.Set("uid", claims.User.ID)
				return next(c)
			}
			return &echo.HTTPError{
				Code:    middleware.ErrJWTInvalid.Code,
				Message: middleware.ErrJWTInvalid.Message,
			}
		}
	}
}

// MustUser 普通用户中间件
func MustUser() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{configure(), user()}
}

// MustAdmin 管理员中间件
func MustAdmin() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{configure(), administrator()}
}
