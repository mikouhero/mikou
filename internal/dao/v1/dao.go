package v1

import (
	//"github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}
