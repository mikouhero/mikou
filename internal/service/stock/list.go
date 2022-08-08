package stock

import (
	"fmt"
)

// 当天股票的所有信息
var StockListURL = "http://16.push2.eastmoney.com/api/qt/clist/get?cb=json&pn=1&pz=1&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&fid=f3&fs=m:0+t:6,m:0+t:80,m:1+t:2,m:1+t:23,m:0+t:81+s:2048&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f22,f11,f62,f128,f136,f115,f152&_=1650533239815"

type List struct {
	Rc   int  `json:"rc"`
	Rt   int  `json:"rt"`
	Svr  int  `json:"svr"`
	Lt   int  `json:"lt"`
	Full int  `json:"full"`
	Data data `json:"data"`
}

type ListData struct {
	Total int            `json:"total"`
	Diff  []ListDataDiff `json:"diff"`
}

type ListDataDiff struct {
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
	F12  float64 `json:"f12"` //代码
	F13  float64 `json:"f13"` //
	F14  string  `json:"f14"` //中文名
	F15  float64 `json:"f15"` //最高价
	F16  float64 `json:"f16"` //最低价
	F17  float64 `json:"f17"` //开盘价
	F18  float64 `json:"f18"` // 昨收
	F19  float64 `json:"f19"`
	F20  float64 `json:"f20"`
	F21  float64 `json:"f21"`
	F22  float64 `json:"f22"`
	F23  float64 `json:"f23"` //市净率
	F24  float64 `json:"f24"`
	F25  float64 `json:"f25"` //今年涨幅
	F62  float64 `json:"f62"`
	F115 float64 `json:"f115"` //市盈率TTM
	F128 float64 `json:"f128"`
	F140 float64 `json:"F140"`
	F141 float64 `json:"f141"`
	F136 float64 `json:"f136"`
	F152 float64 `json:"f152"`
}

// 组装map 数据
func ListMapData() {

	s := httpDo(StockListURL)
	s2 := s[5:]
	s3 := s2[:len(s2)-2]
	fmt.Println(s3)
}
