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
	dsnMap := dsn(*databaseSetting)
	db1, ok := dsnMap["default"]
	if !ok {
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
	//_ = db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	//_ = db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	//_ = db.Callback().Delete().Replace("gorm:delete", deleteCallback)
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

func updateTimeStampForCreateCallback(db *gorm.DB) {
	if db.Statement.Schema != nil {
		nowTime := time.Now()
		if createTimeField, ok := db.Statement.Schema.FieldsByName["CreatedAt"]; ok {

			_ = createTimeField.Set(db.Statement.ReflectValue, nowTime)

		}

		if createTimeField, ok := db.Statement.Schema.FieldsByName["UpdatedAt"]; ok {

			_ = createTimeField.Set(db.Statement.ReflectValue, nowTime)

		}
	}

}

//
//func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
//	if _, ok := scope.Get("gorm:update_column"); !ok {
//		_ = scope.SetColumn("UpdatedAt", time.Now())
//	}
//}

//func deleteCallback(db *gorm.DB) {
//	if !scope.HasError() {
//		var extraOption string
//		if str, ok := scope.Get("gorm:delete_option"); ok {
//			extraOption = fmt.Sprint(str)
//		}
//
//		DeletedAtField, hasDeletedAtField := scope.FieldByName("DeletedAt")
//		if !scope.Search.Unscoped && hasDeletedAtField {
//			now := time.Now()
//			scope.Raw(fmt.Sprintf(
//				"UPDATE %v SET %v=%v%v%v",
//				scope.QuotedTableName(),
//				scope.Quote(DeletedAtField.DBName),
//				scope.AddToVars(now),
//				addExtraSpaceIfExist(scope.CombinedConditionSql()),
//				addExtraSpaceIfExist(extraOption),
//			)).Exec()
//		} else {
//			scope.Raw(fmt.Sprintf(
//				"DELETE FROM %v%v%v",
//				scope.QuotedTableName(),
//				addExtraSpaceIfExist(scope.CombinedConditionSql()),
//				addExtraSpaceIfExist(extraOption),
//			)).Exec()
//		}
//	}
//}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
