//+build wireinject
//go:generate wire

package main

import (
	"github.com/helper-xie/goTemplate/internal/config"
	"github.com/helper-xie/goTemplate/internal/service/example"
	"github.com/helper-xie/goTemplate/internal/svc"

	"github.com/google/wire"
)

func initServ(p string) *example.Service {
	panic(wire.Build(config.Init, svc.NewServiceContext, example.NewExampleService))
}
