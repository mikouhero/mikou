package v1

import (
	"context"
	"mikou/global"
	dao "mikou/internal/dao/v1"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	//svc.dao = dao.New(global.DBEngine)
	svc.dao = dao.New(global.DBEngineV2)
	return svc
}
