package global

import (
	"github.com/jinzhu/gorm"
	gormv2 "gorm.io/gorm"
)

var (
	DBEngine *gorm.DB

	// DBEngine  gorm v2 版本
	DBEngineV2 *gormv2.DB
)
