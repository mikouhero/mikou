package stock

import (
	"encoding/json"
	"fmt"
	"mikou/global"
	v1 "mikou/internal/model/v1"
	"time"
)

// 东财返回结构
type StockList struct {
	Rc   int           `json:"rc"`
	Rt   int           `json:"rt"`
	Svr  int           `json:"svr"`
	Lt   int           `json:"lt"`
	Full int           `json:"full"`
	Data StockListData `json:"data"`
}

// data 数据结构
type StockListData struct {
	Total int                 `json:"total"`
	Diff  []StockListDataDiff `json:"diff"`
}

// 主要数组列数据
type StockListDataDiff struct {
	F1   float64 `json:"f1"`  //
	F2   float64 `json:"f2"`  // 最新价格
	F3   float64 `json:"f3"`  // 	涨跌幅
	F4   float64 `json:"f4"`  //涨跌额
	F5   float64 `json:"f5"`  //成交量
	F6   float64 `json:"f6"`  // 成交额
	F7   float64 `json:"f7"`  //振幅
	F8   float64 `json:"f8"`  // 换手率
	F9   float64 `json:"f9"`  //市盈率
	F10  float64 `json:"f10"` //量比
	F11  float64 `json:"f11"`
	F12  string  `json:"f12"` //代码
	F13  float64 `json:"f13"` //
	F14  string  `json:"f14"` //中文名
	F15  float64 `json:"f15"` //最高价
	F16  float64 `json:"f16"` //最低价
	F17  float64 `json:"f17"` //开盘价
	F18  float64 `json:"f18"` // 昨收
	F19  float64 `json:"f19"`
	F20  float64 `json:"f20"` // 总市值
	F21  float64 `json:"f21"` // 流通市值
	F22  float64 `json:"f22"`
	F23  float64 `json:"f23"` //市净率
	F24  float64 `json:"f24"`
	F25  float64 `json:"f25"`  //今年涨幅
	F26  string  `json:"f26"`  //上市日期
	F35  float64 `json:"f35"`  //内盘
	F62  float64 `json:"f62"`  // 主力净流入
	F184 float64 `json:"f184"` // 主力净占比
	F66  float64 `json:"f66"`  // 超大单净流入
	F69  float64 `json:"f69"`  // 超大单净流入占比
	F72  float64 `json:"f72"`  // 大单净流入
	F75  float64 `json:"f75"`  // 大单净流入占比
	F78  float64 `json:"f78"`  // 中单净流入
	F81  float64 `json:"f81"`  // 中单净流入占比
	F84  float64 `json:"f84"`  // 小单净流入
	F87  float64 `json:"f87"`  // 小单净流入占比
	F115 float64 `json:"f115"` //市盈率TTM
	F128 float64 `json:"f128"`
	F140 float64 `json:"F140"`
	F141 float64 `json:"f141"`
	F136 float64 `json:"f136"`
	F152 float64 `json:"f152"`
	F102 string  `json:"f102"` //地域板块
	F103 string  `json:"f103"` // 概念板块
	F378 float64 `json:"f378"` // 历史最高
}

// 获取所股票列表的url
func getStockListUrl() string {
	cb := "json"
	pn := "1"     // 页码
	pz := "10000" //个数
	po := "1"
	np := "1"
	ut := "bd1d9ddb04089700cf9c27f6f7426281"
	fltt := "2"
	invt := "invt"
	fid := "f3"
	fs := "m:0+t:6,m:0+t:80,m:1+t:2,m:1+t:23,m:0+t:81+s:2048"
	fields := "f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f22,f11,f62,f128,f136,f115,f152"

	return fmt.Sprintf("http://16.push2.eastmoney.com/api/qt/clist/get?cb=%s&pn=%s&pz=%s&po=%s&np=%s&ut=%s&fltt=%s&invt=%s&fid=%s&fs=%s&fields=%s&_=%d",
		cb, pn, pz, po, np, ut, fltt, invt, fid, fs, fields,
		time.Now().Unix(),
	)
}

// 获取所有的股票信息
func GetAllStockListByDFCF() []StockListDataDiff {
	s := httpDo(getStockListUrl())
	s2 := s[5:]
	s3 := s2[:len(s2)-2]
	var a StockList
	_ = json.Unmarshal([]byte(s3), &a)
	return a.Data.Diff
}

// 保存所有code的基本信息
func SaveAllStockListCode() {
	data := GetAllStockListByDFCF()
	var a []map[string]interface{}
	for _, v := range data {
		i := map[string]interface{}{
			"code":  v.F12,
			"name":  v.F14,
			"secid": CodeToSecid(v.F12),
		}
		a = append(a, i)
	}
	global.DBEngineV2.Raw("turncate table stock_code_list")
	global.DBEngineV2.Model(&v1.StockCodeList{}).Create(a)
}
