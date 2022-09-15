package initConfig

import (
	"github.com/robfig/cron/v3"
	"mikou/global"
	v1 "mikou/internal/Job/v1"
)

// initCron  初始化job 任务
func initCron() {
	global.Cron = cron.New(cron.WithSeconds())
	global.Cron.Start()
	v1.JobSetUp()
	//defer global.Cron.Stop()
}
