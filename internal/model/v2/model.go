package v2

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"mikou/pkg/setting"
	"time"
)

type Model struct {
	ID        int        `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:null"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:null"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"default:null"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingSV2) (*gorm.DB, error) {
	fmt.Println(databaseSetting)
	dsnMap := dsn(*databaseSetting)
	db1, ok := dsnMap["default"]
	if (!ok) {
		return nil, errors.New("not set default database")
	}

	db, err := gorm.Open(mysql.Open(db1), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	for key, value := range dsnMap {
		if key == "default" {
			continue
		}
		_ = db.Use(dbresolver.Register(dbresolver.Config{
			Sources: []gorm.Dialector{mysql.Open(value)},
		}, key))
	}

	fmt.Printf("%#v \n", db)

	return db, err
}
func dsn(databaseSetting setting.DatabaseSettingSV2) map[string]string {
	var a = make(map[string]string)
	for key, value := range databaseSetting {
		a[key] = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			value.UserName,
			value.Password,
			value.Host,
			value.DBName,
			value.Charset,
			value.ParseTime,
		)
	}

	return a
}
