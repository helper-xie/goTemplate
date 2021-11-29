package svc

import "github.com/helper-xie/goTemplate/internal/config"

type Ctx struct {
	Cfg *config.Config
}

func NewServiceContext(cfg *config.Config) *Ctx {
	return &Ctx{Cfg: cfg}
}
