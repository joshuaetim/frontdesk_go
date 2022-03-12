package model

import "gorm.io/gorm"

type Visitor struct {
	gorm.Model
	Firstname	string	`json:"first_name" binding:"required"`
	Lastname	string	`json:"last_name" binding:"required"`
	Email		string	`json:"email" binding:"required,email" gorm:"unique"`
	Phone		string	`json:"phone"`
	Occupation	string	`json:"occupation"`
	Note		string	`json:"note"`
	StaffID		uint
	Staff		Staff
	UserID		uint
	User		User
}

func (Visitor) TableName() string {
	return "visitors"
}