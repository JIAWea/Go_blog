package vm

import (
	"log"

	"blog/model"
)

type ResetPasswordRequestViewModel struct {
	LoginViewModel
}

type ResetPasswordRequestViewModelOp struct {}

func (ResetPasswordRequestViewModelOp) GetVM() ResetPasswordRequestViewModel {
	v := ResetPasswordRequestViewModel{}
	v.SetTitle("找回密码")
	return v
}

// 检查邮箱
func CheckEmailExist(email string) bool {
	_, err := model.GetUserByEmail(email)
	if err != nil {
		log.Println("找不到该邮箱:", email)
		return false
	}
	return true
}