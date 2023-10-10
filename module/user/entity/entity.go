package entity

type UserDTO struct {
	ID       uint
	Email    string
	Name     string
	Password string
}

type (
	UserServiceInterface interface {
		CreateUser(req UserDTO) error
		GetAllUser() ([]UserDTO, error)
	}

	UserRepositoryInterface interface {
		InsertUser(userDTO UserDTO) error
		GetAllUser() ([]UserDTO, error)
	}
)
