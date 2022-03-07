package model

import "gorm.io/gorm"

type Staff struct {
	gorm.Model
	Firstname	string	`json:"first_name" binding:"required"`
	Lastname	string	`json:"last_name" binding:"required"`
	Email		string	`json:"email" binding:"required,email" gorm:"unique"`
	Phone		string	`json:"phone"`
	Address		string	`json:"address"`
	Status		string	`json:"status"`
	UserID		int
	User		User
	Visitors	[]Visitor
}

func (Staff) TableName() string {
	return "staff"
}