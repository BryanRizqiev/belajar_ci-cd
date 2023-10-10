package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDBSQLite() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
}

// func InitMigrationSQLite(db *gorm.DB) error {
// 	return db.AutoMigrate(
// 		model.User{},
// 	)
// }
