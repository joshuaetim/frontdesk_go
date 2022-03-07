package repository

import "github.com/joshuaetim/frontdesk/domain/model"

type StaffRepository interface {
	AddStaff(model.Staff) (model.Staff, error)
	GetStaff(int) (model.Staff, error)
	GetAllStaff()([]model.Staff, error)
	UpdateStaff(model.Staff) (model.Staff, error)
	DeleteStaff(model.Staff) error
}