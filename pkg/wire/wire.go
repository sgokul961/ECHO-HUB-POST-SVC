//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/sgokul961/echo-hub-post-svc/pkg/api"
	"github.com/sgokul961/echo-hub-post-svc/pkg/api/handler"
	"github.com/sgokul961/echo-hub-post-svc/pkg/client"
	"github.com/sgokul961/echo-hub-post-svc/pkg/config"
	"github.com/sgokul961/echo-hub-post-svc/pkg/db"
	"github.com/sgokul961/echo-hub-post-svc/pkg/repository"
	"github.com/sgokul961/echo-hub-post-svc/pkg/usecase"
)

func InitApi(cfg config.Config) (*api.ServerHTTP, error) {
	wire.Build(db.Init,
		repository.NewPostRepo,
		usecase.NewPostUseCase,
		handler.NewPostHandler,
		client.NewAuthServiceClient,
		api.NewServerHttp)
	return &api.ServerHTTP{}, nil
}
