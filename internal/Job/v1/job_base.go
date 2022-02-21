package v1

import (
	"github.com/robfig/cron/v3"
	"mikou/global"
	dao "mikou/internal/dao/v1"
	"mikou/pkg/app"
	"time"
)

// 定义一个job 的核心
type JobCore struct {
	Id      int
	Target  string
	Name    string
	EntryId int
	Cron    string
	Args    string
}

// http 类型job
type HttpJob struct {
	JobCore
}

// 函数类型 job
type ExecJob struct {
	JobCore
}

//函数类型的job 执行
func (e *ExecJob) Run() {
	startTime := time.Now()
	// 通过全局job 集合获得指定的执行函数
	var obj = jobList[e.Target]
	if obj == nil {
		global.LoggerV2.Warnf("[Job]  %s  is nil ", e.Target)
		return
	}
	// 执行 job struct 的 Exec 方法
	err := CallExec(obj.(JobsExec), e.Args)
	if err != nil {
		// todo
		global.LoggerV2.Errorf("[Job]  %s  exec failed  %s ", e.Target, err)
	}
	endTime := time.Now()
	spend := endTime.Sub(startTime)
	global.LoggerV2.Infof("[Job] JobCore %s exec success , spend :%v", e.Name, spend)
	return
}

func (h *HttpJob) Run() {
	// todo
}

// 启动时,加载所有job
func JobSetUp() {
	// 初始化job list
	initList()

	d := dao.New(global.DBEngineV2)

	list, err := d.AllJob()
	if err != nil {
		global.LoggerV2.Errorf(" got job list failed %s", err)
		return
	}
	if len(list) == 0 {
		global.LoggerV2.Warn("  job list  is empty")
		return
	}
	for i := 0; i < len(list); i++ {

		if list[i].Status == 0 {
			//忽略禁用的
			continue
		}

		if list[i].Type == 0 {
			//todo

		} else if list[i].Type == 1 {
			job := &ExecJob{JobCore{
				Target: list[i].Target,
				Name:   list[i].Name,
				Id:     list[i].ID,
				Cron:   list[i].Cron,
				Args:   list[i].Args,
			}}
			entryId, _ := AddJob(global.Cron, job)
			_ = d.UpdateJob(list[i].ID, map[string]interface{}{
				"entry_id": entryId,
			})
			job.Run()
		}

	}

}

// 添加job
func AddJob(c *cron.Cron, job Job) (int, error) {
	if job == nil {
		return 0, nil
	}
	return job.addJob(c)
}

// 第三方包添加job
func (job *ExecJob) addJob(c *cron.Cron) (int, error) {
	id, e := c.AddJob(job.Cron, job)
	return int(id), e
}

func RemoveJob(c *cron.Cron, entryID int) chan struct{} {
	ch := make(chan struct{})

	f := func() {
		c.Remove(cron.EntryID(entryID))
		global.LoggerV2.Infof(" [INFO] JobCore Remove success ,info entryID : %d", entryID)
		ch <- struct{}{}
	}
	app.Go(f)
	return ch
}
