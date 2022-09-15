package v1

import (
	"mikou/global"
	v1 "mikou/internal/Job/v1"
	model "mikou/internal/model/v1"
	validate "mikou/internal/validate/v1"
	"mikou/pkg/app"
)

type JobService struct {
	Service
}

// CountJob 统计数量
func (s *Service) CountJob(param *validate.CountJobRequest) (int, error) {

	return s.dao.CountJob(param)
}

//  ListJob 列表数据
func (s *Service) ListJob(param *validate.ListJobRequest, pager *app.Pager) ([]*model.Job, error) {

	return s.dao.ListJob(param, pager.Page, pager.PageSize)
}

// CreateJob  创建job
func (s *Service) CreateJob(param *validate.CreateJobRequest) (int, error) {

	id, err := s.dao.CreateJob(param)
	if err != nil {
		return 0, err
	}
	s.StartJob(id)
	return id, err
}

//  CreateJob  更新信息
func (s *Service) UpdateJob(param *validate.UpdateJobRequest) error {
	var p = make(map[string]interface{})
	if param.Name != "" {
		p["name"] = param.Name
	}
	if param.Args != "" {
		p["args"] = param.Args
	}
	if param.Cron != "" {
		p["cron"] = param.Cron
	}
	if param.Target != "" {
		p["target"] = param.Target
	}

	p["status"] = param.Status
	p["group"] = param.Group
	p["strategy"] = param.Strategy
	p["concurrent"] = param.Concurrent

	err := s.dao.UpdateJob(param.Id, p)
	if err != nil {
		return err
	}
	s.RemoveJob(param.Id)
	s.StartJob(param.Id)
	return nil
}

// 删除信息
func (s *Service) DeleteJob(param *validate.DeleteJobRequest) error {
	s.RemoveJob(param.Id)
	return s.dao.DeleteJob(param.Id)
}

// GetJobByID 获取指定job 信息
func (s *Service) GetJobByID(id int) *model.Job {
	var p = &validate.CommonJob{
		Id: id,
	}
	user, e := s.dao.FindJob(p)

	if e != nil {
		return nil
	}
	return user
}

// StartJob 开启一个job
func (s *Service) StartJob(id int) {
	i := s.GetJobByID(id)
	if i == nil {
		return
	}
	if i.Status == 0 {
		return
	}
	switch i.Type {
	case 0:
		s.startHttpFunc(i)
	case 1:
		s.startFuncJob(i)
	default:
		//todo
	}
}

//startFuncJob  开始一个函数型 job
func (s *Service) startFuncJob(i *model.Job) {
	job := &v1.ExecJob{v1.JobCore{
		Target: i.Target,
		Name:   i.Name,
		Id:     i.ID,
		Cron:   i.Cron,
		Args:   i.Args,
	}}

	entryId, _ := v1.AddJob(global.Cron, job)
	_ = s.dao.UpdateJob(i.ID, map[string]interface{}{
		"entry_id": entryId,
	})
	job.Run()
	return
}

// startHttpFunc 开始一个http型 job
func (s *Service) startHttpFunc(i *model.Job) {
	//todo
	return
}

// RemoveJob 移除一个job
func (s *Service) RemoveJob(id int) {
	i := s.GetJobByID(id)
	if i == nil {
		return
	}
	v1.RemoveJob(global.Cron, i.EntryId)
	_ = s.dao.UpdateJob(i.ID, map[string]interface{}{
		"entry_id": 0,
	})
	return
}
