package response

import "belajar-go-echo/module/user/entity"

type (
	CreateUserResponse struct {
		Message string `json:"message"`
	}
	GetUsersResponse struct {
		Message string           `json:"email"`
		Data    []entity.UserDTO `json:"data"`
	}
)
