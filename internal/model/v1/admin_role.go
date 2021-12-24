package v1

import (
	"gorm.io/gorm"

)

type AdminRole struct {
	Model
	Permission string `json:"permission"`
	Name       string `json:"name"`
}

func (AdminRole) TableName() string {
	return "admin_role_v2"
}

//统计数量
func (a *AdminRole) Count(db *gorm.DB) (int, error) {
	var count int64
	//todo 筛选功能
	if a.Name != "" {
		db = db.Where("name = ? ", a.Name)
	}

	if err := db.Model(&a).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

// 列表
func (a *AdminRole) List(db *gorm.DB, pageOffset, pageSize int) ([]*AdminRole, error) {
	var AdminRoles []*AdminRole
	var err error

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	if a.Name != "" {
		db = db.Where("name = ? ", a.Name)
	}

	if err = db.Find(&AdminRoles).Error; err != nil {
		return nil, err
	}
	return AdminRoles, nil
}

func (a *AdminRole) All(db *gorm.DB) (m []*AdminRole, err error) {
	if err = db.Find(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}
//创建
func (a *AdminRole) Create(db *gorm.DB) (int, error) {
	err := db.Create(&a).Error
	return a.ID, err
}

//更新
func (a *AdminRole) Update(db *gorm.DB, values interface{}) error {
	return db.Model(a).Updates(values).Error
}

// 删除
func (a *AdminRole) Delete(db *gorm.DB) error {
	return db.Delete(&a).Error
}

// 通过模型检索对象
func (a *AdminRole) Find(db *gorm.DB) (*AdminRole, error) {
	e := db.Where(a).First(a).Error
	if e != nil {
		return nil, e
	}
	return a, nil
}
