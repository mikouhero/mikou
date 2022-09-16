package v1

import "github.com/robfig/cron/v3"

//Job 定义 job 接口 类型
type Job interface {
	Run()                                //每个job  必须实现公共 Run 方法
	addJob(cron *cron.Cron) (int, error) // 内部方法 调用第三方库实现
}

// JobsExec  定义一个job 执行结构体类型
type JobsExec interface {
	Exec(arg interface{}) error // Exec()  新增的Job 必须实现的方法，后续统一调用
}

//JobCore 定义一个job 结构体
type JobCore struct {
	Id      int
	Target  string
	Name    string
	EntryId int
	Cron    string
	Args    string
}

// HttpJob  接口类型job
type HttpJob struct {
	JobCore
}

// ExecJob 函数类型job
type ExecJob struct {
	JobCore
}
