package infrastructure

import (
	"github.com/joshuaetim/frontdesk/domain/model"
	"github.com/joshuaetim/frontdesk/domain/repository"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB 
}

func NewUserRepository() repository.UserRepository {
	return &userRepo{
		db: DB(),
	}
}

var _ repository.UserRepository = &userRepo{}

func (r *userRepo) AddUser(user model.User) (model.User, error) {
	return user, r.db.Create(&user).Error
}

func (r *userRepo) GetUser(id int) (model.User, error) {
	var user model.User
	return user, r.db.First(&user, id).Error
}

func (r *userRepo) GetByEmail(email string) (model.User, error) {
	var user model.User
	return user, r.db.First(&user, "email = ?", email).Error
}

func (r *userRepo) GetAllUser() ([]model.User, error) {
	var users []model.User
	return users, r.db.Find(&users).Error
}

func (r *userRepo) UpdateUser(user model.User) (model.User, error) {
	// check if exists
	if err := r.db.First(&user).Error; err != nil {
		return user, err
	}
	return user, r.db.Model(&user).Updates(&user).Error
}

func (r *userRepo) DeleteUser(user model.User) error {
	// exists?
	if err := r.db.First(&user).Error; err != nil {
		return err
	}
	return r.db.Delete(&user).Error
}