package service

import (
	"context"
	"github.com/jiebutech/log"
	"github.com/jiebutech/zin/app/entity"
	"github.com/jiebutech/zin/app/repo"
	"go.uber.org/zap"
)

type DemoSvc struct {
	demoRepo repo.DemoRepo
}

func NewDemoSvc(demoRepo repo.DemoRepo) *DemoSvc {
	return &DemoSvc{demoRepo: demoRepo}
}

func (s DemoSvc) GetById(ctx context.Context, id int64) (*entity.Demo, error) {
	info, err := s.demoRepo.GetDemoByID(id)
	if err != nil {
		log.Error(ctx, "get info error", zap.Error(err))
		return nil, err
	}
	return info, nil
}
