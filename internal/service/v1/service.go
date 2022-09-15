package v1

import (
	"context"
	"mikou/global"
	dao "mikou/internal/dao/v1"
)

// Service  服务层 controller 调用层
type Service struct {
	ctx context.Context // 上下文
	dao *dao.Dao        // dao 数据层
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	//svc.dao = dao.New(global.DBEngine)
	svc.dao = dao.New(global.DBEngineV2)
	return svc
}
