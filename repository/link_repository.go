package repository

import (
	"gorm.io/gorm"
	"shortlink/model"
)

type LinkRepositoryInterface interface {
	FindByName(name string) (*model.Link, error)
	Create(name string, url string, userID uint) (*model.Link, error)
}

type linkRepository struct {
	db *gorm.DB
}

func NewLinkRepository(db *gorm.DB) LinkRepositoryInterface {
	return &linkRepository{db}
}

func (ur *linkRepository) FindByName(name string) (link *model.Link, err error) {
	err = ur.db.Where("name = ?", name).First(&link).Error
	return
}

func (ur *linkRepository) Create(name string, url string, userID uint) (user *model.Link, err error) {
	user = &model.Link{
		Name: name,
		Url:  url,
		UserID: userID,
	}

	if err := ur.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
