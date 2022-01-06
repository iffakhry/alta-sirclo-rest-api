package datastore

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	connectionString := "root:Alch3mist@tcp(127.0.0.1:3306)/db_sirclo_api_gorm?charset=utf8&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(connectionString), &gorm.Config{})
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
