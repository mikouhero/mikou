package v1

import (
	//"github.com/jinzhu/gorm"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type AdminMenu struct {
	Model
	Pid        int    `json:"pid"`
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	Urls       string `json:"urls"`
	Remark     string `json:"remark"`
	Sort       int    `json:"sort"`
	Status     int    `json:"status"`
	Button     int    `json:"Button"`
	UniqueFlag string `json:"unique_flag"`
	Check      int    `json:"check"`
}

func (AdminMenu) TableName() string {
	return "admin_menu_v2"
}

func (a *AdminMenu) All(db *gorm.DB) (m []*AdminMenu, err error) {
	if err = db.Clauses(dbresolver.Use("center")).Find(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

//创建
func (a *AdminMenu) Create(db *gorm.DB) (int, error) {
	err := db.Create(&a).Error
	return a.ID, err
}

//更新
func (a *AdminMenu) Update(db *gorm.DB, values interface{}) error {
	return db.Model(a).Updates(values).Error
}

// 删除
func (a *AdminMenu) Delete(db *gorm.DB) error {
	return db.Delete(&a).Error
}

// 通过模型检索对象
func (a *AdminMenu) Find(db *gorm.DB) (*AdminMenu, error) {
	e := db.Where(a).First(a).Error
	if e != nil {
		return nil, e
	}
	return a, nil
}

func (a *AdminMenu) GetListByIDs(db *gorm.DB, ids []int) (m []*AdminMenu, err error) {
	db = db.Order("pid asc,sort asc")
	if ids == nil {
		err = db.Find(&m).Error
	} else {
		err = db.Or("`check` = ?",0).Find(&m, ids).Error
	}
	return m, err
}
