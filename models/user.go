package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
