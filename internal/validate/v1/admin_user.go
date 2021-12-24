package v1

type CountAdminUserRequest = ListAdminUserRequest

type ListAdminUserRequest struct {
	Account string `form:"account" json:"account" binding:"max=100"`
	Name    string `form:"name" json:"name" binding:"max=100"`
	RoleId  int    `form:"role_id" json:"role_id"`
	Status  int    `form:"status" json:"status" binding:"oneof=-1 0 1"`
}

type CreateAdminUserRequest struct {
	Account  string `form:"account" json:"account" binding:"required,max=100"`
	Name     string `form:"name" json:"name" binding:"required,max=100"`
	Password string `form:"password" json:"password" binding:"required,max=100"`
	RoleId   int    `form:"role_id" json:"role_id" binding:"required"`
	Status   int    `form:"status" json:"status" binding:"required,oneof= 0 1"`
	Avatar   string `form:"avatar" json:"avatar"`
}

type UpdateAdminUserRequest struct {
	Id       int    `form:"id" json:"id" binding:"required"`
	Account  string `form:"account" json:"account" binding:"max=100"`
	Name     string `form:"name" json:"name" binding:"max=100"`
	Password string `form:"password" json:"password" binding:"max=100"`
	RoleId   int    `form:"role_id" json:"role_id" `
	Status   int    `form:"status" json:"status" binding:"oneof=0 1"`
	Avatar   string `form:"avatar" json:"avatar"`
}

type DeleteAdminUserRequest struct {
	Id int `form:"id" json:"id" binding:"required"`
}

type CommonAdminUser struct {
	Id       int
	Account  string
	Name     string
	Password string
}
