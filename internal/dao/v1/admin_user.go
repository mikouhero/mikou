package v1

import (
	model "mikou/internal/model/v1"
	validate "mikou/internal/validate/v1"
	"mikou/pkg/app"
)

// 统计数量
func (d *Dao) CountAdminUser(param *validate.CountAdminUserRequest) (int, error) {

	adminUser := &model.AdminUser{
		Account: param.Account,
		Name:    param.Name,
		RoleId:  param.RoleId,
		//Status:  param.Status,
	}
	return adminUser.Count(d.engine)
}

func (d *Dao) ListAdminUser(param *validate.ListAdminUserRequest, page, pageSize int) ([]*model.AdminUserForDao, error) {
	adminUser := &model.AdminUser{
		Account: param.Account,
		Name:    param.Name,
		RoleId:  param.RoleId,
		//Status:  param.Status,
	}
	pageOffset := app.GetPageOffset(page, pageSize)

	return adminUser.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateAdminUser(param *validate.CreateAdminUserRequest) (int, error) {
	adminUser := &model.AdminUser{
		Account:  param.Account,
		Name:     param.Name,
		Password: param.Password,
		RoleId:   param.RoleId,
		Status:   param.Status,
		Avatar:   param.Avatar,
	}
	return adminUser.Create(d.engine)

}

func (d *Dao) UpdateAdminUser(id int, values map[string]interface{}) error {
	adminUser := model.AdminUser{
		Model: model.Model{ID: id},
	}
	return adminUser.Update(d.engine, values)
}

func (d *Dao) DeleteAdminUser(id int) error {
	adminUser := model.AdminUser{
		Model: model.Model{ID: id},
	}
	return adminUser.Delete(d.engine)
}

func (d *Dao) FindAdminUser(param *validate.CommonAdminUser) (*model.AdminUserForDao, error) {

	user := model.AdminUser{
		Model:    model.Model{ID: param.Id},
		Name:     param.Name,
		Account:  param.Account,
		Password: param.Password,
	}
	return user.Find(d.engine)
}
