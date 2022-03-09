package repository

import "github.com/joshuaetim/frontdesk/domain/model"

type VisitorRepository interface {
	AddVisitor(model.Visitor) (model.Visitor, error)
	GetVisitor(int) (model.Visitor, error)
	GetAllVisitor() ([]model.Visitor, error)
	UpdateVisitor(model.Visitor) (model.Visitor, error)
	DeleteVisitor(model.Visitor) error
}