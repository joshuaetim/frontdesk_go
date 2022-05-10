package repository

import "github.com/joshuaetim/frontdesk/domain/model"

type UserRepository interface {
	AddUser(model.User) (model.User, error)
	GetUser(uint) (model.User, error)
	GetByEmail(string) (model.User, error)
	GetAllUser() ([]model.User, error)
	UpdateUser(model.User) (model.User, error)
	DeleteUser(model.User) error
	GetUserStaff(uint) ([]model.Staff, error)
	GetUserVisitors(uint) ([]model.Visitor, error)
	CountUsers() int
}
