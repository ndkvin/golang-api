package service

import (
	"shortlink/repository"
)

type AuthServiceInterface interface {
	register() error
}

type authService struct {
	userRepository repository.UserRepositoryInterface
}

func NewAuthService(userRepository repository.UserRepositoryInterface) AuthServiceInterface {
	return &authService{userRepository}
}

func (a *authService) register() error {
	return nil
}