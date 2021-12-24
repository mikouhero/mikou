package v1

type LoginAdminLoginRequest struct {
	Account  string `form:"account" json:"account" binding:"required,max=100"`
	Password string `form:"password" json:"password" binding:"required,max=100"`
}
