package controller

import (
	"belajar-go-echo/module/auth/controller/request"
	"belajar-go-echo/module/auth/controller/response"
	"belajar-go-echo/module/auth/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	AuthService entity.AuthServiceInterface
}

func New(authService entity.AuthServiceInterface) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

func (this *AuthController) Login(ctx echo.Context) error {
	req := new(request.LoginRequest)
	var token string
	var err error
	if err = ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.LoginResponse{
			Message: "Request not valid",
		})
	}
	token, err = this.AuthService.Login(req.Email, req.Password)
	if err != nil {
		if err.Error() == "record not found" {
			return ctx.JSON(http.StatusBadRequest, response.LoginResponse{
				Message: "Invalid credentials",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, response.LoginResponse{
			Message: "Server error",
		})
	}
	return ctx.JSON(http.StatusOK, response.LoginResponse{
		Message: "Success login",
		Token:   token,
	})
}
