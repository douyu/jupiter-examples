//go:build wireinject
// +build wireinject

package service

import (
	"github.com/douyu/jupiter-examples/uuid/internal/pkg/redis"
	"github.com/google/wire"
)

func createMockUuidService() *Uuid {
	panic(wire.Build(
		NewUuidService,
		redis.ProviderSet,
		wire.Struct(new(Options), "*"),
	))
}
