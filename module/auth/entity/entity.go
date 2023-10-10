package entity

import (
	userEntity "belajar-go-echo/module/user/entity"

	"github.com/labstack/echo/v4"
)

type (
	AuthServiceInterface interface {
		Login(email, password string) (string, error)
		ExtractToken(e echo.Context) (uint, string)
	}
	AuthRepositoryInterface interface {
		CheckUser(email, password string) (userEntity.UserDTO, error)
	}
)
