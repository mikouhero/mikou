package stock

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var host string = "https://push2his.eastmoney.com/api/qt/stock/kline/get"

type klineS struct {
	Rc   int  `json:"rc"`
	Rt   int  `json:"rt"`
	Svr  int  `json:"svr"`
	Lt   int  `json:"lt"`
	Full int  `json:"full"`
	Data data `json:"data"`
}
type data struct {
	Code      string  `json:"code"`
	Market    int     `json:"market"`
	Name      string  `json:"name"`
	Decimal   int     `json:"decimal"`
	Dktotal   int     `json:"dktotal"`
	PreKPrice float64 `json:"pre_k_price"`
	Klines    []string
}

// 获取历史记录接口

/**
https://push2his.eastmoney.com/api/qt/stock/kline/get?cb=json&fields1=f1,f2,f3,f4,f5,f6&fields2=f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61&ut=7eea3edcaed734bea9cbfc24409ed989&klt=101&fqt=1&secid=0.000815&beg=0&end=20500000

cb  返回格式 json
fields1 f1,f2,f3,f4,f5,f6
fields2   'f51': '日期',
        'f52': '开盘',
        'f53': '收盘',
        'f54': '最高',
        'f55': '最低',
        'f56': '成交量',
        'f57': '成交额',
        'f58': '振幅',
        'f59': '涨跌幅',
        'f60': '涨跌额',
        'f61': '换手率',

   # 沪市指数
    if rawcode[:3] == '000':
        return f'1.{rawcode}'
    # 深证指数
    if rawcode[:3] == '399':
        return f'0.{rawcode}'
    # 沪市股票
    if rawcode[0] != '6':
        return f'0.{rawcode}'
    # 深市股票
    return f'1.{rawcode}'



  klt: k线间距 默认为 101 即日k
            klt:1 1 分钟
            klt:5 5 分钟
            klt:101 日
            klt:102 周

   fqt: 复权方式
            不复权 : 0
            前复权 : 1
            后复权 : 2
*/
/**



http://push2.eastmoney.com/api/qt/kamt/get?fields1=f1,f2,f3,f4&fields2=f51,f52,f53,f54&ut=b2884a393a59ad64002292a3e90d46a5&cb=jQuery112408500394100136952_1647336428422&_=1647336428682


http://push2his.eastmoney.com/api/qt/stock/kline/get?cb=jQuery112408500394100136952_1647336428552&fields1=f1,f2,f3,f4,f5,f6&fields2=f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61&ut=7eea3edcaed734bea9cbfc24409ed989&klt=102&fqt=1&secid=0.000815&beg=0&end=20500000&_=1647336428710
*/

func CodeToSecid(code string) string {
	//// 沪市指数
	//if code[:3] == "000" {
	//	return "1." + code
	//}
	//// 深证指数
	//if code[:3] == "399" {
	//	return "0." + code
	//}
	// 沪市股票
	if code[0:1] != "6" {
		return "0." + code
	}
	// 深市股票
	return "1." + code

}

func GetDayKline(Secid string) []map[string]interface{} {
	//gotData(Secid, 20220316, 20220316)
	return gotData(Secid, 0, 0)

}

func DayKlineUrl(Secid string, start int, end int) string {

	if start == 0 && end == 0 {
		end = 20500000
	}
	return fmt.Sprintf(host+"?cb=%s&fields1=%s&fields2=%s&klt=%s&fqt=%s&secid=%s&beg=%d&end=%d&_=%d",
		"json",
		"f1,f2,f3,f4,f5,f6",
		"f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61",
		"101",
		"1",
		Secid,
		start,
		end,
		time.Now().Unix(),
	)
}

func mapData(a []string) []map[string]interface{} {
	var data []map[string]interface{}

	for _, v := range a {
		arr := strings.Split(v, `,`)
		i := map[string]interface{}{
			"date":              arr[0],
			"begin":             arr[1],
			"end":               arr[2],
			"highest":           arr[3],
			"lowest":            arr[4],
			"turnover":          arr[5],
			"amount":            arr[6],
			"amplitude":         arr[7],
			"fluctuation_range": arr[8],
			"amount_range":      arr[9],
			"turnover_rate":     arr[10],
		}
		data = append(data, i)
	}
	return data
}

func gotData(Secid string, start int, end int) []map[string]interface{} {

	url := DayKlineUrl(Secid, start, end)
	s := httpDo(url)
	s2 := s[5:]
	s3 := s2[:len(s2)-2]
	var a klineS
	_ = json.Unmarshal([]byte(s3), &a)
	return mapData(a.Data.Klines)
}

func httpDo(url string) string {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// handle error
	}

	// 添加请求头
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; Touch; rv:11.0) like Gecko")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	req.Header.Add("Referer", "http://quote.eastmoney.com/center/gridlist.html")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Println("err")
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	return string(b)
}
