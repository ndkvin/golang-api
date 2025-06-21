package service

import (
	"errors"
	"shortlink/dto/auth"
	customError "shortlink/error"
	"shortlink/model"
	"shortlink/repository"
	"shortlink/util"

	"gorm.io/gorm"
)

type AuthServiceInterface interface {
	Register(req auth.RegisterRequest) (*model.User, error)
	Login(req auth.LoginRequest) (string, error)
}

type authService struct {
	userRepository repository.UserRepositoryInterface
}

func NewAuthService(userRepository repository.UserRepositoryInterface) AuthServiceInterface {
	return &authService{userRepository}
}

func (a *authService) Register(req auth.RegisterRequest) (user *model.User, err error) {
	user, err = a.userRepository.FindByEmail(req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if user.ID != 0 {
		return nil, &customError.BadRequest{Message: "Email already exists"}
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user, err = a.userRepository.Create(req.Email, hashedPassword)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *authService) Login(req auth.LoginRequest) (token string, err error) {
	user, err := a.userRepository.FindByEmail(req.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", &customError.NotFound{Message: "User not found"}
	}

	if !util.CheckPasswordHash(req.Password, user.Password) {
		return "", &customError.BadRequest{Message: "Invalid password"}
	}

	token, err = util.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
