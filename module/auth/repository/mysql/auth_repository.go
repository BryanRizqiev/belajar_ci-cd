package mysql_auth_repository

import (
	authEntity "belajar-go-echo/module/auth/entity"
	userEntity "belajar-go-echo/module/user/entity"
	"belajar-go-echo/module/user/repository/model"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) authEntity.AuthRepositoryInterface {
	return &AuthRepository{
		db: db,
	}
}

func (this *AuthRepository) CheckUser(email string, password string) (userEntity.UserDTO, error) {
	var user model.User
	var userDTO userEntity.UserDTO
	tx := this.db.Where("email = ?", email).Where("password = ?", password).First(&user)
	if tx.Error != nil {
		return userDTO, tx.Error
	}
	userDTO = userEntity.UserDTO{
		ID:    user.ID,
		Email: user.Email,
		Name:  *user.Name,
	}
	return userDTO, nil
}
