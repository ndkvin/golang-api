package service

import (
	"errors"
	"shortlink/dto/link"
	customError "shortlink/error"
	"shortlink/model"
	"shortlink/repository"

	"gorm.io/gorm"
)

type LinkServiceInterface interface {
	Create(req link.CreateLinkRequest, userID uint) (*model.Link, error)
}

type linkService struct {
	linkRepository repository.LinkRepositoryInterface
}

func NewLinkService(linkRepository repository.LinkRepositoryInterface) LinkServiceInterface {
	return &linkService{linkRepository}
}

func (l *linkService) Create(req link.CreateLinkRequest, userID uint) (link *model.Link, err error) {
	link, err = l.linkRepository.FindByName(req.Name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if link.ID != 0 {
		return nil, &customError.BadRequest{Message: "Name Already Taken"}
	}


	link, err = l.linkRepository.Create(req.Name, req.Url, userID)
	if err != nil {
		return nil, err
	}

	return link, nil
}