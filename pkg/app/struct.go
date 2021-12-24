package app

import (
	"encoding/json"
)

//使用json tag，转换结构体
func SourceToTarget(sourceStruct interface{}, targetStruct interface{}) {
	marshal, _ := json.Marshal(sourceStruct)
	json.Unmarshal(marshal, targetStruct)

	return
}
