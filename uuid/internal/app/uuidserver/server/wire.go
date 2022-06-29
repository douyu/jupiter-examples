//go:build wireinject
// +build wireinject

package server

import (
	"github.com/douyu/jupiter"
	"github.com/douyu/jupiter-examples/uuid/internal/app/uuidserver/controller"
	"github.com/douyu/jupiter-examples/uuid/internal/app/uuidserver/service"
	"github.com/google/wire"
)

func InitApp(app *jupiter.Application) error {
	panic(wire.Build(
		wire.Struct(new(Options), "*"),
		controller.ProviderSet,
		service.ProviderSet,
		ProviderSet,
		initApp,
	))
}
