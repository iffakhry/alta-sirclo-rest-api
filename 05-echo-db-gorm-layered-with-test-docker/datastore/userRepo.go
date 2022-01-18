package datastore

import (
	"fmt"

	"gorm.io/gorm"
)

type Identity struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}

func GetUsers(db *gorm.DB) ([]UserResponse, error) {
	// db.Query("SELECT id, name, email, address from users")
	var users []UserResponse
	err := db.Find(&users).Error
	//for
	// result.scan
	// append ke users
	return users, err
}

func AddUser(db *gorm.DB, user *User) error {
	return db.Save(user).Error
}

func GetUserByIdentity(db *gorm.DB, identity Identity) (User, error) {
	var user User
	if err := db.Where(identity).Find(&user).Error; err != nil {
		// terjadi error waktu query
		return user, err
	}
	if user.Name == identity.Name {
		// usernya benar-benar ada
		return user, nil
	}
	// tidak error, tapi usernya tidak ada
	return user, fmt.Errorf("user not found")
}
