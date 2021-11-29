package example

import "github.com/helper-xie/goTemplate/internal/svc"

type Service struct {
	ctx *svc.Ctx
}

func NewExampleService(ctx *svc.Ctx) *Service {
	return &Service{ctx: ctx}
}

func (s *Service) Start() error {
	return nil
}
