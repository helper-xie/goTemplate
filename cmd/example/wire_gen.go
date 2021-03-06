// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/helper-xie/goTemplate/internal/config"
	"github.com/helper-xie/goTemplate/internal/service/example"
	"github.com/helper-xie/goTemplate/internal/svc"
)

// Injectors from wire.go:

func initServ(p string) *example.Service {
	configConfig := config.Init(p)
	ctx := svc.NewServiceContext(configConfig)
	service := example.NewExampleService(ctx)
	return service
}
