package model

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model

	Email    string
	Name     *string
	Password string
}
