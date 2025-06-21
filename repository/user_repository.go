package repository 

import (
	"shortlink/model"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	FindByEmail(email string) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &userRepository{db}
}

func (ur *userRepository) FindByEmail(email string) (user model.User, err error) {
	err = ur.db.Where("email = ?", email).First(&user).Error
	return
}