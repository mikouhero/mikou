package v1

import (
	"github.com/jinzhu/gorm"
	model "mikou/internal/model/v1"
	validate "mikou/internal/validate/v1"
	"mikou/pkg/app"
)

// 统计数量
func (d *Dao) CountAdminRole(param *validate.CountAdminRoleRequest) (int, error) {

	AdminRole := &model.AdminRole{
		Name: param.Name,
	}
	return AdminRole.Count(d.engine)
}

func (d *Dao) AllAdminRole() (m []*model.AdminRole, err error) {
	AdminRole := &model.AdminRole{}
	return AdminRole.All(d.engine)

}

func (d *Dao) ListAdminRole(param *validate.ListAdminRoleRequest, page, pageSize int) ([]*model.AdminRole, error) {
	AdminRole := &model.AdminRole{
		Name: param.Name,
	}
	pageOffset := app.GetPageOffset(page, pageSize)

	return AdminRole.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateAdminRole(param *validate.CreateAdminRoleRequest) (int, error) {
	AdminRole := &model.AdminRole{
		Name:       param.Name,
		Permission: param.Permission,
	}
	return AdminRole.Create(d.engine)

}

func (d *Dao) UpdateAdminRole(id int, values map[string]interface{}) error {
	AdminRole := model.AdminRole{
		Model: model.Model{ID: id},
	}
	return AdminRole.Update(d.engine, values)
}

func (d *Dao) DeleteAdminRole(id int) error {
	AdminRole := model.AdminRole{
		Model: model.Model{ID: id},
	}
	return AdminRole.Delete(d.engine)
}

func (d *Dao) FindAdminRole(id int) (*model.AdminRole, error) {
	role := model.AdminRole{
		Model: model.Model{ID: id},
	}

	adminRole, e := role.Find(d.engine)
	if e == gorm.ErrRecordNotFound {
		e = nil
	}
	return adminRole, e
}
