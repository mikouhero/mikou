package v1

import (
	"gorm.io/gorm"
)

type StockDayKline struct {
	Model
	Date              string `json:"date"`
	Begin             float64
	End               float64
	Highest           float64
	Lowest            float64
	Turnover          float64
	Amount            float64
	Amplitude         float64
	Fluctuation_range float64
	Amount_range      float64
	Turnover_rate     float64
	T_5               float64
	T_10              float64
	T_20              float64
	T_60              float64
	T_120             float64
	D_5               float64
	D_10              float64
	D_20              float64
	D_60              float64
	D_120             float64
	Tr_5              float64
	Tr_10             float64
}

func (StockDayKline) TableName() string {
	return "stock_day_kline"
}

func (a *StockDayKline) CreateBatch(db *gorm.DB, b []*StockDayKline) error {
	return db.CreateInBatches(b, 100).Error
}

func (a *StockDayKline) All(db *gorm.DB) (m []*StockDayKline, err error) {
	if err = db.Find(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (a *StockDayKline) Update(db *gorm.DB, values interface{}) error {
	return db.Model(a).Updates(values).Error
}
