package list

import (
	"fmt"
	"mikou/global"
	v1 "mikou/internal/model/v1"
	"mikou/internal/service/stock"
	"strconv"
)

// 新添加的job 必须按照以下格式定义，并实现Exec函数
type StockDayKline struct {
}

// 执行函数
func (t StockDayKline) Exec(arg interface{}) error {
	// TODO: 这里需要注意 Examples 传入参数是 string 所以 arg.(string)；请根据对应的类型进行转化；

	data := stock.GetDayKline("0.000815")
	global.DBEngineV2.Model(&v1.StockDayKline{}).Create(data)
	return nil
}

var all []*v1.StockDayKline

func Handel() {
	s := &v1.StockDayKline{}
	all, _ = s.All(global.DBEngineV2)

	for k, v := range all {
		if v.D_5 == 0 {
			updateData(k, 5)
		}
		if v.D_10 == 0 {
			updateData(k, 10)
		}
		if v.D_20 == 0 {
			updateData(k, 20)
		}

		if v.D_60 == 0 {
			updateData(k, 60)
		}

		if v.D_120 == 0 {
			updateData(k, 120)
		}
	}
}
func updateData(k int, day int) {
	if k < day-1 {
		return
	}
	klines := all[k-(day-1) : k+1]
	var D5 float64
	var T5 float64
	var TR5 float64
	for _, kline := range klines {
		D5 += kline.End
		T5 += kline.Turnover
		TR5 += kline.Turnover_rate

	}
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", D5/float64(day)), 64)
	value1, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", T5/float64(day)), 64)
	value2, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", TR5/float64(day)), 64)

	var p = make(map[string]interface{})
	days := strconv.Itoa(day)
	p["D_"+days] = value
	p["T_"+days] = value1
	p["TR_"+days] = value2

	s := v1.StockDayKline{
		Model: v1.Model{ID: all[k].ID},
	}
	s.Update(global.DBEngineV2, p)
}
