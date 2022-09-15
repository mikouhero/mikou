package v1

import (
	model "mikou/internal/model/v1"
	validate "mikou/internal/validate/v1"
	"mikou/pkg/app"
)

// CountJob   统计job统计数量
func (d *Dao) CountJob(param *validate.CountJobRequest) (int, error) {

	Job := &model.Job{
		Name:   param.Name,
		Status: param.Status,
		Group:  param.Group,
	}
	return Job.Count(d.engine)
}

// ListJob  job 列表
func (d *Dao) ListJob(param *validate.ListJobRequest, page, pageSize int) ([]*model.Job, error) {
	Job := &model.Job{
		Name:   param.Name,
		Status: param.Status,
		Group:  param.Group,
	}
	pageOffset := app.GetPageOffset(page, pageSize)

	return Job.List(d.engine, pageOffset, pageSize)
}

// CreateJob  创建job
func (d *Dao) CreateJob(param *validate.CreateJobRequest) (int, error) {
	Job := &model.Job{
		Name:       param.Name,
		Status:     param.Status,
		Group:      param.Group,
		Type:       param.Type,
		Cron:       param.Cron,
		Target:     param.Target,
		Args:       param.Args,
		Strategy:   param.Strategy,
		Concurrent: param.Concurrent,
	}
	return Job.Create(d.engine)

}

// UpdateJob 更新 JOb
func (d *Dao) UpdateJob(id int, values map[string]interface{}) error {
	Job := model.Job{
		Model: model.Model{ID: id},
	}
	return Job.Update(d.engine, values)
}

// DeleteJob  删除 JOb
func (d *Dao) DeleteJob(id int) error {
	Job := model.Job{
		Model: model.Model{ID: id},
	}
	return Job.Delete(d.engine)
}

// FindJob 查询指定 JOb 得信息
func (d *Dao) FindJob(param *validate.CommonJob) (*model.Job, error) {

	job := model.Job{
		Model: model.Model{ID: param.Id},
	}
	return job.Find(d.engine)
}

// AllJob  获取所以的job
func (d *Dao) AllJob() (m []*model.Job, err error) {
	Job := &model.Job{}
	return Job.All(d.engine)

}
