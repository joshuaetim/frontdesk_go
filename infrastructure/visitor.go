package infrastructure

import (
	"github.com/joshuaetim/frontdesk/domain/model"
	"github.com/joshuaetim/frontdesk/domain/repository"
	"gorm.io/gorm"
)

type visitorRepo struct {
	db *gorm.DB
}

func NewVisitorRepository() repository.VisitorRepository {
	return &visitorRepo{
		db: DB(),
	}
}

var _ repository.VisitorRepository = &visitorRepo{}

func (r *visitorRepo) AddVisitor(visitor model.Visitor) (model.Visitor, error) {
	return visitor, r.db.Create(&visitor).Error
}

func (r *visitorRepo) GetVisitor(id int) (model.Visitor, error) {
	var visitor model.Visitor
	return visitor, r.db.First(&visitor, id).Error
}

func (r *visitorRepo) GetAllVisitor() ([]model.Visitor, error) {
	var visitors []model.Visitor
	return visitors, r.db.Find(&visitors).Error
}

func (r *visitorRepo) UpdateVisitor(visitor model.Visitor) (model.Visitor, error) {
	// exists?
	if err := r.db.First(&visitor).Error; err != nil {
		return visitor, err
	}
	return visitor, r.db.Model(&visitor).Updates(&visitor).Error
}

func (r *visitorRepo) DeleteVisitor(visitor model.Visitor) error {
	// exists?
	if err := r.db.First(&visitor).Error; err != nil {
		return err
	}
	return r.db.Delete(&visitor).Error
}