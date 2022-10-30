package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"uniqueIndex"`
	Password  string
}

type Token struct {
	gorm.Model
	Token string `gorm:"uniqueIndex"`
	UserID int
	User User
}
