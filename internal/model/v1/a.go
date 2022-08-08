package v1

import (
	"gorm.io/gorm"
)

type A struct {
	Name       string `json:"name" gorm:"size:255;"`     // 名称
	Group      int    `json:"group" gorm:"size:1;"`      // 任务分组
	Type       int    `json:"type" gorm:"size:1;"`       // 任务类型
	Cron       string `json:"cron" gorm:"size:255;"`     // cron表达式
	Target     string `json:"target" gorm:"size:255;"`   // 调用目标
	Args       string `json:"args" gorm:"size:255;"`     // 目标参数
	Strategy   int    `json:"strategy" gorm:"size:255;"` // 执行策略
	Concurrent int    `json:"concurrent" gorm:"size:1;"` // 是否并发
	Status     int    `json:"status" gorm:"size:1;"`     // 状态
	EntryId    int    `json:"entry_id" gorm:"size:11;"`  // job启动时返回的id
	Model
}

func (a A) TableName() string {
	return a.Name + "_table"
}

func (a *A) All(db *gorm.DB) (m []*Job, err error) {
	if err = db.Find(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}
