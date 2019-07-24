// Copyright 2018 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:generate protoc -I proto --micro_out=casbinpb --go_out=casbinpb casbinpb/casbin.proto

package main

import (
	"log"
	"net/http"
	"os"

	metrics "github.com/ProtocolONE/go-micro-plugins/wrapper/monitoring/prometheus"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/client/selector/static"
	casbinpb "github.com/unistack-org/casbin-micro/casbinpb"
)

func main() {
	options := []micro.Option{
		micro.Name("casbin"),
		micro.Version("0.0.1"),
		micro.WrapHandler(metrics.NewHandlerWrapper()),
		micro.AfterStop(func() error {
			app.logger.Info("Micro service stopped")
			app.Stop()
			return nil
		}),
	}

	if os.Getenv("MICRO_SELECTOR") == "static" {
		log.Println("Use micro selector `static`")
		options = append(options, micro.Selector(static.NewSelector()))
	}

	app.logger.Info("Initialize micro service")

	app.service = micro.NewService(options...)

	if err := app.svc.Init(); err != nil {
		app.logger.Fatal("Create service instance failed", zap.Error(err))
	}

	err = casbinpb.RegisterCasbinHandler(app.service.Server(), app.svc)

	if err != nil {
		app.logger.Fatal("Service init failed", zap.Error(err))
	}

	app.router = http.NewServeMux()
	app.initHealth()
	app.initMetrics()
}
