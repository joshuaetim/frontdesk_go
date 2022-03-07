package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Firstname	string	`json:"first_name" binding:"required"`
	Lastname	string	`json:"last_name" binding:"required"`
	Email		string	`json:"email" binding:"required,email" gorm:"unique"`
	Password	string	`json:"password"`
	Company		string	`json:"company"`
	Staff		[]Staff
	Visitors	[]Visitor
}

func (u *User) PublicUser() *User {
	return &User{
		Firstname: u.Firstname,
		Lastname: u.Lastname,
		Email: u.Email,
		Company: u.Company,
	}
}

// table name
func (User) TableName() string {
	return "users"
}