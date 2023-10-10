package database

import (
	"belajar-go-echo/app/config"
	"belajar-go-echo/module/user/repository/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMysql(cfg *config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func InitMigrationMySQL(db *gorm.DB) error {
	db.Migrator().DropTable(model.User{})
	return db.AutoMigrate(
		model.User{},
	)
}
