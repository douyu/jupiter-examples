// Copyright 2020 Douyu
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"time"

	"github.com/douyu/jupiter"
	"github.com/douyu/jupiter/pkg/client/redis"
	"github.com/douyu/jupiter/pkg/xlog"
)

// run: go run main.go -config=config.toml
type Engine struct {
	jupiter.Application
}

func NewEngine() *Engine {
	eng := &Engine{}
	if err := eng.Startup(
		eng.exampleForRedis,
		eng.exampleForRedisStub,
		eng.exampleForRedisClusterStub,
	); err != nil {
		xlog.Default().Panic("startup", xlog.Any("err", err))
	}
	return eng
}

func main() {
	app := NewEngine()
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func (eng *Engine) exampleForRedisStub() (err error) {
	//build redisStub
	redisStub := redis.StdConfig("myredis").MustSingleton()
	// set string
	setRes := redisStub.CmdOnMaster().Set(context.Background(), "jupiter-redis", "redisStub", time.Second*5)
	xlog.Default().Info("redisStub set string", xlog.Any("res", setRes))
	// get string
	getRes := redisStub.CmdOnMaster().Get(context.Background(), "jupiter-redis")
	xlog.Default().Info("redisStub get string", xlog.Any("res", getRes))
	return
}
func (eng *Engine) exampleForRedisClusterStub() (err error) {
	//build redisClusterStub
	redisStub := redis.StdConfig("myredis").MustSingleton()
	// set string
	setRes := redisStub.CmdOnMaster().Set(context.Background(), "jupiter-redisCluster", "redisClusterStub", time.Second*5)
	xlog.Default().Info("redisClusterStub set string", xlog.Any("res", setRes))
	// get string
	getRes := redisStub.CmdOnMaster().Get(context.Background(), "jupiter-redisCluster")
	xlog.Default().Info("redisClusterStub get string", xlog.Any("res", getRes))
	return
}

func (eng *Engine) exampleForRedis() (err error) {
	//build redisStub
	redisClient := redis.StdConfig("myredistub").MustSingleton()
	// set string
	setRes := redisClient.CmdOnMaster().Set(context.Background(), "jupiter-redis", "redisStub", time.Second*5)
	xlog.Default().Info("redisStub set string", xlog.Any("res", setRes))
	// get string
	getRes := redisClient.CmdOnMaster().Get(context.Background(), "jupiter-redis")
	xlog.Default().Info("redisStub get string", xlog.Any("res", getRes))

	//build redisClusterStub
	redisClient = redis.StdConfig("myrediscluster").MustSingleton()
	// set string
	setRes = redisClient.CmdOnMaster().Set(context.Background(), "jupiter-redisCluster", "redisClusterStub", time.Second*5)
	xlog.Default().Info("redisClusterStub set string", xlog.Any("res", setRes))
	// get string
	getRes = redisClient.CmdOnMaster().Get(context.Background(), "jupiter-redisCluster")
	xlog.Default().Info("redisClusterStub get string", xlog.Any("res", getRes))
	return
}
