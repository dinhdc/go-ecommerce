//go:build wireinject

package wire

import (
	"github.com/dinhdc/go-ecommerce/internal/controllers"
	"github.com/dinhdc/go-ecommerce/internal/repo"
	"github.com/dinhdc/go-ecommerce/internal/service"
	"github.com/google/wire"
)

func InitUserRouterHandler() (*controllers.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controllers.NewUserController,
	)
	return new(controllers.UserController), nil
}
