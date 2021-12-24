package v1

import (
	v1 "mikou/internal/model/v1"
	validate "mikou/internal/validate/v1"
	"mikou/pkg/app"
)

func (s *Service) AllAdminMenu() ([]TreeList, error) {
	menus, e := s.dao.AllAdminMenu()

	if e != nil || menus == nil {
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

func (s *Service) CreateAdminMenu(param *validate.CreateAdminMenuRequest) (int, error) {
	return s.dao.CreateAdminMenu(param)
}

func (s *Service) UpdateAdminMenu(param *validate.UpdateAdminMenuRequest) error {
	var p = make(map[string]interface{})
	if param.Name != "" {
		p["name"] = param.Name
	}
	if param.Name != "" {
		p["Icon"] = param.Icon
	}
	if param.Name != "" {
		p["Urls"] = param.Urls
	}
	if param.Name != "" {
		p["Remark"] = param.Remark
	}
	p["status"] = param.Status
	p["button"] = param.Button
	p["check"] = param.Check
	p["sort"] = param.Sort
	p["pid"] = param.Pid

	return s.dao.UpdateAdminMenu(param.Id, p)
}

func (s *Service) DeleteAdminMenu(param *validate.DeleteAdminMenuRequest) error {
	return s.dao.DeleteAdminMenu(param.Id)
}

func (s *Service) FindAdminMenu(param *validate.FindAdminMenuRequest) (*v1.AdminMenu, error) {
	return s.dao.FindAdminMenu(param.Id)
}

type Menu struct {
	ID         int    `json:"id"`
	Pid        int    `json:"pid"`
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	Urls       string `json:"urls"`
	Remark     string `json:"remark"`
	Sort       int    `json:"sort"`
	Status     int    `json:"status"`
	UniqueFlag string `json:"unique_flag"`
	Check      int    `json:"check"`
	Button     int    `json:"button"`
}

// menu 树状结构
type TreeList struct {
	Menu
	Children []TreeList `json:"children"`
}

//
func GetMenuTree(menuList []Menu, pid int) []TreeList {
	treeList := []TreeList{}
	for _, v := range menuList {
		if v.Pid == pid {
			child := GetMenuTree(menuList, v.ID)
			node := TreeList{
				Menu: Menu{
					ID:         v.ID,
					Pid:        v.Pid,
					Name:       v.Name,
					Icon:       v.Icon,
					Urls:       v.Urls,
					Remark:     v.Remark,
					Sort:       v.Sort,
					Status:     v.Status,
					UniqueFlag: v.UniqueFlag,
					Check:      v.Check,
					Button:     v.Button,
				},
			}
			node.Children = child
			treeList = append(treeList, node)
		}
	}
	return treeList
}
