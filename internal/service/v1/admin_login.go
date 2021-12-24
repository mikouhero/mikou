package v1

import (
	validate "mikou/internal/validate/v1"
	"mikou/pkg/app"
	"mikou/pkg/errcode"
)

func (s *Service) AdminUserLogin(param validate.LoginAdminLoginRequest) (string, error) {
	user := s.GetUserByAccountAndPwd(param.Account, param.Password)

	if user == nil {
		return "", errcode.AdminUserAccountOrPwdWrong
	}
	if user.Status != 1 {
		return "", errcode.AdminUserAccountBan
	}
	token, e := app.GenerateToken(user.Account, user.RoleId, user.ID)
	if e != nil {
		return "", e
	}
	return token, nil

}
