package datastore

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
}

type UserResponse struct {
	ID      string `json:"id" form:"id"`
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Address string `json:"address" form:"address"`
}
