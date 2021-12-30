package v1

import (
	"fmt"
	//"github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

type AdminUser struct {
	Model
	Account  string `json:"account"`
	Name     string `json:"name"`
	Password string `json:"-"`
	RoleId   int    `json:"role_id"`
	Status   int    `json:"status"`
	Avatar   string `json:"avatar"`
	Token    string `json:"token"`
}

type AdminUserForDao struct {
	AdminUser
	RoleName string `json:"role_name"`
}

func (AdminUser) TableName() string {
	return "admin_user_v2"
}

//统计数量
func (a *AdminUser) Count(db *gorm.DB) (int, error) {
	var count int64
	//todo 筛选功能
	if a.Account != "" {
		db = db.Where("account = ? ", a.Account)
	}

	if a.Name != "" {
		db = db.Where("name = ? ", a.Name)
	}

	if a.RoleId != 0 {
		db = db.Where("role_id = ? ", a.RoleId)
	}

	//if a.Status != -1 {
	//	db = db.Where("status = ? ", a.Status)
	//}

	if err := db.Model(&a).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

// 用户列表
func (a *AdminUser) List(db *gorm.DB, pageOffset, pageSize int) ([]*AdminUserForDao, error) {
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	if a.Account != "" {
		db = db.Where("account = ? ", a.Account)
	}

	if a.Name != "" {
		db = db.Where("name = ? ", a.Name)
	}

	if a.RoleId != 0 {
		db = db.Where("role_id = ? ", a.RoleId)
	}
	//if a.Status != -1 {
	//	db = db.Where("status = ? ", a.Status)
	//}
	var res []*AdminUserForDao
	err = db.Table(a.TableName()).
		Select("admin_user_v2.*,p2.name as role_name").
		Joins("left join admin_role_v2 as p2 on p2.id = admin_user_v2.role_id").
		Where("admin_user_v2.deleted_at IS NULL").
		Scan(&res).Error

	return res, err
}

//创建用户
func (a *AdminUser) Create(db *gorm.DB) (int, error) {
	fmt.Println(a)
	err := db.Create(&a).Error
	return a.ID, err
}

//更新用户信息
func (a *AdminUser) Update(db *gorm.DB, values interface{}) error {
	return db.Model(a).Updates(values).Error
}

// 删除
func (a *AdminUser) Delete(db *gorm.DB) error {
	return db.Delete(&a).Error
}

// 通过模型检索对象
func (a *AdminUser) Find(db *gorm.DB) (*AdminUserForDao, error) {

	var res []*AdminUserForDao
	e := db.Table(a.TableName()).
		Where(a).
		Select("admin_user_v2.*,p2.name as role_name").
		Joins("left join admin_role_v2 as p2 on p2.id = admin_user_v2.role_id").
		Scan(&res).Error
	//e := db.Where(a).First(a).Error
	if e != nil || len(res) == 0 {
		return nil, e
	}
	return res[0], nil
}
