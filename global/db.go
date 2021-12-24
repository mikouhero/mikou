package global

import (
	"github.com/jinzhu/gorm"
	gormv2 "gorm.io/gorm"
)

var (
	DBEngine *gorm.DB
	DBEngineV2 *gormv2.DB
)
