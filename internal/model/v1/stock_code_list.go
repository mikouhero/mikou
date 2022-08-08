package v1

import "gorm.io/gorm"

// 所有股票的代码集合
type StockCodeList struct {
	Model
	Code  string `json:"code"`  // 数字代码
	Name  string `json:"name"`  //  中文名
	Secid string `json:"secid"` // 东财 code代码
}

// 数据表表名
func (StockCodeList) TableName() string {
	return "stock_code_list"
}

// 返回表中所有数据
func (a *StockCodeList) All(db *gorm.DB) (m []*StockCodeList, err error) {
	if err = db.Find(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}
