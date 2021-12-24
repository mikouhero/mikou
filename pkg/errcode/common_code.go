package errcode

var (
	Success         = NewError(0, "成功")
	ServerError     = NewError(500, "服务内部错误")
	InvalidParams   = NewError(400, "入参错误")
	TooManyRequests = NewError(440, "请求过多")
	Forbidden       = NewError(403, "没有权限访问")

	AdminUserAccountHadExist   = NewError(10000, "后台账号已存在")
	AdminUserAccountOrPwdWrong = NewError(10001, "账号或者密码错误")
	AdminUserAccountBan        = NewError(10002, "账号禁用中")
	UnauthorizedTokenTimeout   = NewError(10003, "token 过期")
	UnauthorizedTokenError     = NewError(10004, "token错误")
)
