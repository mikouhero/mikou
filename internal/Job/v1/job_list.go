package v1

import "mikou/internal/Job/v1/list"

// 所有函数类型  执行job 的集合
var jobList map[string]JobsExec

func initList() {
	jobList = map[string]JobsExec{
		"ExamplesOne": list.ExamplesOne{},
	}
}