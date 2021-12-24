package v1

type CountAdminRoleRequest = ListAdminRoleRequest

type ListAdminRoleRequest struct {
	Name string `form:"name" json:"name" binding:"max=100"`
}

type CreateAdminRoleRequest struct {
	Name       string `form:"name" json:"name" binding:"required,max=100"`
	Permission string `form:"permission" json:"permission" binding:"required"`
}

type UpdateAdminRoleRequest struct {
	Id         int    `form:"id" json:"id" binding:"required"`
	Name       string `form:"name" json:"name" binding:"required,max=100"`
	Permission string `form:"permission" json:"permission" binding:"required"`
}

type DeleteAdminRoleRequest struct {
	Id int `form:"id" json:"id" binding:"required"`
}


type FindAdminRoleRequest struct {
	Id int `form:"id" json:"id" binding:"required"`
}