//go:build wireinject
// +build wireinject

package main

import (
	"shortlink/database"
	"shortlink/handler"
	"shortlink/repository"
	"shortlink/service"

	"github.com/google/wire"
)

func InitializeWire() (*handler.Handler, error) {
	wire.Build(
		database.NewDatabase,

		repository.NewUserRepository,
		repository.NewLinkRepository,

		service.NewAuthService,
		service.NewLinkService,

		handler.NewAuthHandler,
		handler.NewLinkHandler,
		handler.NewHandler,
	)
	return &handler.Handler{}, nil
}