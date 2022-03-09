package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Firstname	string	`json:"first_name"`
	Lastname	string	`json:"last_name"`
	Email		string	`json:"email" binding:"required,email" gorm:"unique"`
	Password	string	`json:"password" binding:"required"`
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