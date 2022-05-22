//go:build wireinject
// +build wireinject

package e2e

import (
	"github.com/google/wire"
	"uuid/internal/app/uuidserver/service"
	"uuid/internal/pkg/redis"
)

func CreateUuidService(redisCli redis.RedisInterface) *service.Uuid {
	panic(wire.Build(
		service.NewUuidService,
		wire.Struct(new(service.Options), "Redis"),
	))
}
