package datastore

import "gorm.io/gorm"

func GetUsers(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Find(&users).Error
	return users, err
}

func AddUser(db *gorm.DB, user *User) error {
	return db.Save(user).Error
}
