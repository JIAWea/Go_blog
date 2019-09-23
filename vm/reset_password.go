package vm

import (
	"blog/model"
)

type ResetPasswordViewModel struct {
	LoginViewModel
	Token string
}

type ResetPasswordViewModelOp struct{}

// GetVM
func (ResetPasswordViewModelOp) GetVM(token string) ResetPasswordViewModel {
	v := ResetPasswordViewModel{}
	v.SetTitle("设置密码")
	v.Token = token
	return v
}

// CheckToken func
func CheckToken(tokenString string) (string, error) {
	return model.CheckToken(tokenString)
}

// ResetUserPassword func
func ResetUserPassword(username, password string) error {
	return model.UpdatePassword(username, password)
}