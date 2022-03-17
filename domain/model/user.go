package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email" binding:"email" gorm:"unique"`
	Password  string `json:"password"`
	Company   string `json:"company"`
	Staff     []Staff
	Visitors  []Visitor
}

func (u *User) PublicUser() *User {
	user := &User{
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Email:     u.Email,
		Company:   u.Company,
	}
	user.ID = u.ID
	user.CreatedAt = u.CreatedAt
	return user
}

// table name
func (User) TableName() string {
	return "users"
}
