// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package service

import (
	"github.com/douyu/jupiter-examples/uuid/internal/pkg/redis"
)

// Injectors from wire.go:

func createMockUuidService() *Uuid {
	redisInterface := redis.NewRedis()
	options := Options{
		Redis: redisInterface,
	}
	uuid := NewUuidService(options)
	return uuid
}
