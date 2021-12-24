package v1

type CreateAdminMenuRequest struct {
	Pid        int    `json:"pid" form:"pid" binding:"min=0"`
	Name       string `json:"name" form:"name" binding:"required,max=100"`
	Icon       string `json:"icon" form:"icon" binding:"max=100"`
	Urls       string `json:"urls" form:"urls" binding:"required,max=100"`
	Remark     string `json:"remark" form:"remark" binding:"max=100"`
	Sort       int    `json:"sort" form:"sort" binding:"min=0"`
	Status     int    `json:"status" form:"status" binding:"oneof=0 1"`
	Button     int    `json:"button" form:"button" binding:"oneof=0 1"`
	Check     int    `json:"check" form:"check" binding:"oneof=0 1"`
	UniqueFlag string `json:"unique_flag" form:"unique_flag" binding:"max=100"`
}

type UpdateAdminMenuRequest struct {
	Id         int    `form:"id" json:"id" binding:"required,min=1"`
	Pid        int    `json:"pid" form:"pid" binding:"min=0"`
	Name       string `json:"name" form:"name" binding:"required,max=100"`
	Icon       string `json:"icon" form:"icon" binding:"max=100"`
	Urls       string `json:"urls" form:"urls" binding:"required,max=100"`
	Remark     string `json:"remark" form:"remark" binding:"max=100"`
	Sort       int    `json:"sort" form:"sort" binding:"min=0"`
	Status     int    `json:"status" form:"status" binding:"oneof=0 1"`
	Button     int    `json:"button" form:"button" binding:"oneof=0 1"`
	Check     int    `json:"check" form:"check" binding:"oneof=0 1"`
	UniqueFlag string `json:"unique_flag" form:"unique_flag" binding:"max=100"`
}

type DeleteAdminMenuRequest struct {
	Id int `form:"id" json:"id" binding:"required,min=1"`
}

type	FindAdminMenuRequest struct {
	Id int `form:"id" json:"id" binding:"required,min=1"`
}
