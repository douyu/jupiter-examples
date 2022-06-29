//go:build wireinject
// +build wireinject

package e2e

import (
	"github.com/douyu/jupiter-examples/uuid/internal/app/uuidserver/service"
	"github.com/douyu/jupiter-examples/uuid/internal/pkg/redis"
	"github.com/google/wire"
)

func CreateUuidService(redisCli redis.RedisInterface) *service.Uuid {
	panic(wire.Build(
		service.NewUuidService,
		wire.Struct(new(service.Options), "Redis"),
	))
}
