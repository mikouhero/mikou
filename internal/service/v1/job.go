package v1

import (
	model "mikou/internal/model/v1"
	validate "mikou/internal/validate/v1"
	"mikou/pkg/app"
)

type JobService struct {
	Service
}

// 统计数量
func (s *Service) CountJob(param *validate.CountJobRequest) (int, error) {

	return s.dao.CountJob(param)
}

// 列表数据
func (s *Service) ListJob(param *validate.ListJobRequest, pager *app.Pager) ([]*model.Job, error) {

	return s.dao.ListJob(param, pager.Page, pager.PageSize)
}

func (s *Service) CreateJob(param *validate.CreateJobRequest) (int, error) {

	return s.dao.CreateJob(param)
}

// 更新信息
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

	return s.dao.UpdateJob(param.Id, p)
}

// 删除信息
func (s *Service) DeleteJob(param *validate.DeleteJobRequest) error {
	return s.dao.DeleteJob(param.Id)
}

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

func StartJob()  {

}

func  RemoveJob()  {
	
}