package main

import (
	"belajar-go-echo/app"
	"belajar-go-echo/app/config"
	"belajar-go-echo/app/database"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	db := database.InitDBMysql(cfg)
	if err := database.InitMigrationMySQL(db); err != nil {
		panic(err)
	}

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] status=${status} method=${method} uri=${uri} latency=${latency_human} ip=${remote_ip}\n",
	}))

	app.InitApp(db, e, config.SECRET_JWT, jwt.SigningMethodHS256)
	e.Logger.Fatal(e.Start(":8080"))
}
