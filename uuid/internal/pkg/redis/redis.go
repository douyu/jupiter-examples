package redis

import (
	"context"
	"strings"

	xredis "github.com/douyu/jupiter/pkg/client/redis"
	"github.com/google/wire"
)

var (
	ProviderSet = wire.NewSet(
		NewRedis,
	)

	redisNodeIdKey = "jupiter.uuid.node"
)

type Redis struct {
	*xredis.Client
}

func NewRedis() RedisInterface {
	return &Redis{
		xredis.StdConfig("uuid").MustSingleton(),
	}
}

// todo 采用 redis 的原子操作，这里目前是为了先实现功能
func (r *Redis) GetNodeId() (int64, error) {
	nodeID, err := r.CmdOnMaster().Get(context.TODO(), redisNodeIdKey).Int64()
	if err != nil && !strings.Contains(err.Error(), "redis: nil") {
		return 0, err
	}

	nodeID++

	r.CmdOnMaster().Set(context.TODO(), redisNodeIdKey, nodeID, 0)

	return nodeID, nil
}
