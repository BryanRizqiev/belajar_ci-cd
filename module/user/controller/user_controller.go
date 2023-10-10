package controller

import (
	"belajar-go-echo/module/user/controller/request"
	"belajar-go-echo/module/user/controller/response"
	"belajar-go-echo/module/user/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService entity.UserServiceInterface
}

func NewUserController(userService entity.UserServiceInterface) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (this *UserController) CreateUserController(ctx echo.Context) error {
	req := new(request.CreateUserRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.CreateUserResponse{
			Message: "Request not valid",
		})
	}
	userDTO := entity.UserDTO{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	}
	if err := this.userService.CreateUser(userDTO); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.CreateUserResponse{
			Message: "Server error",
		})
	}
	return ctx.JSON(http.StatusCreated, response.CreateUserResponse{
		Message: "Success create user",
	})
}

func (this *UserController) GetAllUserController(ctx echo.Context) error {
	var usersDTO []entity.UserDTO
	var err error
	usersDTO, err = this.userService.GetAllUser()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.GetUsersResponse{
			Message: "Server error",
		})
	}
	return ctx.JSON(http.StatusOK, response.GetUsersResponse{
		Message: "Success get users",
		Data:    usersDTO,
	})
}
