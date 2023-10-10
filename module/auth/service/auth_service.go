package service

import (
	"belajar-go-echo/module/auth/entity"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthService struct {
	Secret         string
	SigningMethod  jwt.SigningMethod
	AuthRepository entity.AuthRepositoryInterface
}

func New(secret string, signingMethod jwt.SigningMethod, authRepository entity.AuthRepositoryInterface) entity.AuthServiceInterface {
	return &AuthService{
		Secret:         secret,
		SigningMethod:  signingMethod,
		AuthRepository: authRepository,
	}
}

func (this *AuthService) Login(email, password string) (string, error) {
	userDTO, err := this.AuthRepository.CheckUser(email, password)
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{}
	claims["id"] = userDTO.ID
	claims["name"] = userDTO.Name
	claims["email"] = userDTO.Email
	claims["exp"] = time.Now().Add(30 * time.Minute).Unix()
	token := jwt.NewWithClaims(this.SigningMethod, claims)
	return token.SignedString([]byte(this.Secret))
}

func (this *AuthService) ExtractToken(e echo.Context) (uint, string) {
	panic("unimplemented")
}
