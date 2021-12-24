package v1

import (
	model "mikou/internal/model/v1"
	validate "mikou/internal/validate/v1"
	"mikou/pkg/app"
)

type AdminUserService struct {
	Service
}

// 统计数量
func (s *Service) CountAdminUser(param *validate.CountAdminUserRequest) (int, error) {

	return s.dao.CountAdminUser(param)
}

// 列表数据
func (s *Service) ListAdminUser(param *validate.ListAdminUserRequest, pager *app.Pager) ([]*model.AdminUserForDao, error) {

	return s.dao.ListAdminUser(param, pager.Page, pager.PageSize)
}

// 新增用户信息
func (s *Service) CreateAdminUser(param *validate.CreateAdminUserRequest) (int, error) {

	param.Password = app.MD5(param.Password)
	return s.dao.CreateAdminUser(param)
}

// 更新用户信息
func (s *Service) UpdateAdminUser(param *validate.UpdateAdminUserRequest) error {
	var p = make(map[string]interface{})

	if param.Account !="" {
		p["account"] = param.Account
	}
	if param.Name != "" {
		p["name"] = param.Name
	}
	if param.Avatar != "" {
		p["avatar"] = param.Avatar
	}
	if param.RoleId != 0 {
		p["role_id"] = param.RoleId
	}
	if param.Password != "" {
		//todo
		p["password"] = app.MD5(param.Password)
	}
	p["status"] = param.Status

	return s.dao.UpdateAdminUser(param.Id, p)
}

// 删除用户信息
func (s *Service) DeleteAdminUser(param *validate.DeleteAdminUserRequest) error {
	return s.dao.DeleteAdminUser(param.Id)
}

// 通过账号判断是否存在用户
func (s *Service) CheckAccountExist(account string) *model.AdminUserForDao {
	var p = &validate.CommonAdminUser{
		Account: account,
	}
	user, e := s.dao.FindAdminUser(p)

	if e != nil {
		return nil
	}
	return user

}

//通过account 和pwd 获取用户信息
func (s *Service) GetUserByAccountAndPwd(account, password string) *model.AdminUserForDao {
	var p = &validate.CommonAdminUser{
		Account:  account,
		Password: app.MD5(password),
	}
	user, e := s.dao.FindAdminUser(p)

	if e != nil {
		return nil
	}
	return user
}

func (s *Service) GetUserByID(id int) *model.AdminUserForDao {
	var p = &validate.CommonAdminUser{
		Id:  id,
	}
	user, e := s.dao.FindAdminUser(p)

	if e != nil {
		return nil
	}
	return user
}
