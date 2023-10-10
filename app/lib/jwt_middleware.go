package lib

import (
	"belajar-go-echo/app/config"

	echoJwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echoJwt.WithConfig(echoJwt.Config{
		SigningKey: []byte(config.SECRET_JWT),
	})
}
