package v1

import (
	"encoding/json"
	model "mikou/internal/model/v1"
	validate "mikou/internal/validate/v1"
	"mikou/pkg/app"
)

func (s *Service) CountAdminRole(param *validate.CountAdminRoleRequest) (int, error) {
	return s.dao.CountAdminRole(param)
}

func (s *Service) ListAdminRole(param *validate.ListAdminRoleRequest, pager *app.Pager) ([]*model.AdminRole, error) {
	return s.dao.ListAdminRole(param, pager.Page, pager.PageSize)
}

func (s *Service) AllAdminRole() ([]*model.AdminRole, error) {
	return s.dao.AllAdminRole()
}

func (s *Service) CreateAdminRole(param *validate.CreateAdminRoleRequest) (int, error) {
	return s.dao.CreateAdminRole(param)
}

func (s *Service) UpdateAdminRole(param *validate.UpdateAdminRoleRequest) error {
	var p = make(map[string]interface{})
	if param.Name != "" {
		p["name"] = param.Name
	}
	if param.Permission != "" {
		p["permission"] = param.Permission
	}

	return s.dao.UpdateAdminRole(param.Id, p)
}

func (s *Service) DeleteAdminRole(param *validate.DeleteAdminRoleRequest) error {
	return s.dao.DeleteAdminRole(param.Id)

}
func (s *Service) FindAdminRole(param *validate.FindAdminRoleRequest)(*model.AdminRole, error)  {
	return s.dao.FindAdminRole(param.Id)

}

//  获取角色所有生效菜单
func (s *Service) GetRoleMenuList(id int) ([]*model.AdminMenu, error) {

	role, e := s.dao.FindAdminRole(id)

	if e != nil || role == nil {
		return nil, e
	}

	var ids []int
	if role.Name != "admin" {
		// 权限集合数组
		e = json.Unmarshal([]byte(role.Permission), &ids)
	}
	return s.dao.GetAdminMenuListByIds(ids)

}

//  获取角色所有生效树状菜单
func (s *Service) GetRoleMenuTreeList(id int) ([]TreeList, error) {
	menus, e := s.GetRoleMenuList(id)
	if e != nil {
		return nil, e
	}
	var m []Menu
	for _, value := range menus {
		var a Menu
		app.SourceToTarget(value, &a)
		m = append(m, a)
	}
	tree := GetMenuTree(m, 0)
	return tree, nil
}
