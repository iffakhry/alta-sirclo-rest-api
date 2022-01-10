package datastore

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(connectionString string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(connectionString), &gorm.Config{})
}

func InitialMigration(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
