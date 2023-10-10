package service

import (
	"belajar-go-echo/module/user/entity"
)

type UserService struct {
	userRepository entity.UserRepositoryInterface
}

func NewUserServie(repository entity.UserRepositoryInterface) entity.UserServiceInterface {
	return &UserService{
		userRepository: repository,
	}
}

func (this *UserService) CreateUser(userDTO entity.UserDTO) error {
	if err := this.userRepository.InsertUser(userDTO); err != nil {
		return err
	}
	return nil
}

func (this *UserService) GetAllUser() ([]entity.UserDTO, error) {
	var usersDTO []entity.UserDTO
	var err error
	usersDTO, err = this.userRepository.GetAllUser()
	if err != nil {
		return nil, err
	}
	return usersDTO, nil
}
