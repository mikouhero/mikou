package v1

import "github.com/robfig/cron/v3"

// 定义 job 接口 类型
type Job interface {
	Run()
	addJob(cron *cron.Cron) (int, error)
}

// 定义 执行job 接口类型
type JobsExec interface {
	Exec(arg interface{}) error
}

// 定义执行 job 函数
func CallExec(e JobsExec, arg interface{}) error {
	return e.Exec(arg)
}
