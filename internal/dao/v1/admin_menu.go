package v1

import (
	model "mikou/internal/model/v1"
	validate "mikou/internal/validate/v1"
)

func (d *Dao) AllAdminMenu() (m []*model.AdminMenu, err error) {
	AdminMenu := &model.AdminMenu{}
	return AdminMenu.All(d.engine)

}

func (d *Dao) CreateAdminMenu(param *validate.CreateAdminMenuRequest) (int, error) {
	AdminMenu := &model.AdminMenu{
		Name:       param.Name,
		Pid:        param.Pid,
		Icon:       param.Icon,
		Urls:       param.Urls,
		Remark:     param.Remark,
		Sort:       param.Sort,
		Status:     param.Status,
		Button:     param.Button,
		UniqueFlag: param.UniqueFlag,
		Check:      param.Check,
	}
	return AdminMenu.Create(d.engine)

}

func (d *Dao) UpdateAdminMenu(id int, values map[string]interface{}) error {
	AdminMenu := model.AdminMenu{
		Model: model.Model{ID: id},
	}
	return AdminMenu.Update(d.engine, values)
}

func (d *Dao) DeleteAdminMenu(id int) error {
	AdminMenu := model.AdminMenu{
		Model: model.Model{ID: id},
	}
	return AdminMenu.Delete(d.engine)
}

func (d *Dao) FindAdminMenu(id int) (*model.AdminMenu, error) {
	AdminMenu := model.AdminMenu{
		Model: model.Model{ID: id},
	}
	return AdminMenu.Find(d.engine)
}

func (d *Dao) GetAdminMenuListByIds(ids []int) ([]*model.AdminMenu, error) {
	AdminMenu := model.AdminMenu{}
	return AdminMenu.GetListByIDs(d.engine, ids)
}
