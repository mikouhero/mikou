package list

import (
	"mikou/global"
	"time"
)

var timeFormat = "2006-01-02 15:04:05"

// 新添加的job 必须按照以下格式定义，并实现Exec函数
type ExamplesOne struct {
}

// 执行函数
func (t ExamplesOne) Exec(arg interface{}) error {
	str := time.Now().Format(timeFormat) + " [INFO] JobCore ExamplesOne exec success"
	// TODO: 这里需要注意 Examples 传入参数是 string 所以 arg.(string)；请根据对应的类型进行转化；
	var args string
	switch arg.(type) {

	case string:
		if arg.(string) != "" {
			args = arg.(string)
		} else {
			// todo
		}
		break
	}
	global.LoggerV2.Infof(str + " arg :" + args)
	return nil
}
