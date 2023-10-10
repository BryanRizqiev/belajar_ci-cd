package app

import (
	"belajar-go-echo/app/lib"
	authController "belajar-go-echo/module/auth/controller"
	mysql_auth_repository "belajar-go-echo/module/auth/repository/mysql"
	authService "belajar-go-echo/module/auth/service"
	userController "belajar-go-echo/module/user/controller"
	"belajar-go-echo/module/user/repository/mysql"
	userService "belajar-go-echo/module/user/service"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitApp(db *gorm.DB, e *echo.Echo, jwtSecret string, signingMethod jwt.SigningMethod) {
	userRepository := mysql.NewUserRepository(db)
	userService := userService.NewUserServie(userRepository)
	userController := userController.NewUserController(userService)

	authRepository := mysql_auth_repository.New(db)
	authService := authService.New(jwtSecret, signingMethod, authRepository)
	authController := authController.New(authService)

	e.POST("/login", authController.Login)
	e.POST("/users", userController.CreateUserController)
	e.GET("/users", userController.GetAllUserController, lib.JWTMiddleware())
}
