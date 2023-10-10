package config

import "os"

var (
	SECRET_JWT string
)

type AppConfig struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOSTNAME string
	DB_PORT     string
	DB_NAME     string
	JWT_KEY     string
}

func GetConfig() *AppConfig {
	return Config()
}

func Config() *AppConfig {
	app := AppConfig{}
	app.DB_USERNAME = os.Getenv("DB_USERNAME")
	app.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	app.DB_HOSTNAME = os.Getenv("DB_HOST")
	app.DB_PORT = "3306"
	app.DB_NAME = os.Getenv("DB_NAME")
	app.JWT_KEY = os.Getenv("JWT_SECRET")
	SECRET_JWT = app.JWT_KEY
	return &app
}
