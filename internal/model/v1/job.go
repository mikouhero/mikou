package v1

import (
	"gorm.io/gorm"
)

type Job struct {
	Name       string `json:"name" gorm:"size:255;"`              // 名称
	Group      int    `json:"group" gorm:"size:1;"`               // 任务分组
	Type       int    `json:"type" gorm:"size:1;"`                // 任务类型
	Cron       string `json:"cron" gorm:"size:255;"`              // cron表达式
	Target     string `json:"target" gorm:"size:255;"`            // 调用目标
	Args       string `json:"args" gorm:"size:255;"`              // 目标参数
	Strategy   int    `json:"strategy" gorm:"size:255;"`          // 执行策略
	Concurrent int    `json:"concurrent" gorm:"size:1;"`          // 是否并发
	Status     int    `json:"status" gorm:"size:1;"`              // 状态
	EntryId    int    `json:"entry_id" gorm:"size:11;"`           // job启动时返回的id
	Model
}

func (Job) TableName() string {
	return "job"
}
func (a *Job) Count(db *gorm.DB) (int, error) {
	var count int64
	//todo 筛选功能

	if a.Name != "" {
		db = db.Where("name = ? ", a.Name)
	}

	if a.Status != 0 {
		db = db.Where("status = ? ", a.Status)
	}

	if a.Group != 0 {
		db = db.Where("Group = ? ", a.Group)
	}
	if err := db.Model(&a).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (a *Job) List(db *gorm.DB, pageOffset, pageSize int) ([]*Job, error) {
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if a.Name != "" {
		db = db.Where("name = ? ", a.Name)
	}

	if a.Status != 0 {
		db = db.Where("status = ? ", a.Status)
	}

	if a.Group != 0 {
		db = db.Where("Group = ? ", a.Group)
	}
	var res []*Job
	err = db.Table(a.TableName()).
		Select("*").
		Where("deleted_at IS NULL").
		Scan(&res).Error

	return res, err
}

func (a *Job) Create(db *gorm.DB) (int, error) {
	err := db.Create(&a).Error
	return a.ID, err
}

func (a *Job) Update(db *gorm.DB, values interface{}) error {
	return db.Model(a).Updates(values).Error
}

// 删除
func (a *Job) Delete(db *gorm.DB) error {
	return db.Delete(&a).Error
}

// 通过模型检索对象
func (a *Job) Find(db *gorm.DB) (*Job, error) {

	e := db.Where(a).First(a).Error

	if e != nil {
		return nil, e
	}
	return a, nil
}


func (a *Job) All(db *gorm.DB) (m []*Job, err error) {
	if err = db.Find(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}
