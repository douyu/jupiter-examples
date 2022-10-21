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
	"fmt"
	"math/rand"
	"time"

	"github.com/douyu/jupiter"
	"github.com/douyu/jupiter/pkg/core/sentinel"
	"github.com/douyu/jupiter/pkg/xlog"
)

// run: go run main.go -config=config.toml
type Engine struct {
	jupiter.Application
}

func NewEngine() *Engine {
	eng := &Engine{}
	if err := eng.Startup(
		eng.exampleSentinel,
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

func (eng *Engine) exampleSentinel() (err error) {
	err = sentinel.StdConfig().Build()
	if err != nil {
		panic(fmt.Sprintf("sentinel init failed: %s", err.Error()))
	}

	for k := 0; k < 20; k++ {

		e, b := sentinel.Entry("some-test")
		if b != nil {
			// 请求被拒绝，在此处进行处理
			fmt.Println("Rejected", b.Error())
			time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
		} else {
			// 请求允许通过，此处编写业务逻辑
			fmt.Println("Passed")
			time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)

			// 务必保证业务结束后调用 Exit
			e.Exit()
		}
	}

	return
}
