package mysql

import (
	"belajar-go-echo/module/user/entity"
	"belajar-go-echo/module/user/repository/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) entity.UserRepositoryInterface {
	return &UserRepository{
		db: db,
	}
}

func (this *UserRepository) GetAllUser() ([]entity.UserDTO, error) {
	var users []model.User
	if tx := this.db.Find(&users); tx.Error != nil {
		return nil, tx.Error
	}
	var usersDTO []entity.UserDTO
	for _, user := range users {
		userDTO := entity.UserDTO{
			ID:    user.ID,
			Email: user.Email,
			Name:  *user.Name,
		}
		usersDTO = append(usersDTO, userDTO)
	}
	return usersDTO, nil
}

func (this *UserRepository) InsertUser(userDTO entity.UserDTO) error {
	user := model.User{
		Name:     &userDTO.Name,
		Email:    userDTO.Email,
		Password: userDTO.Password,
	}
	if tx := this.db.Create(&user); tx.Error != nil {
		return tx.Error
	}
	return nil
}
