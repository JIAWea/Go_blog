package vm

import (
	"log"

	"blog/model"
)

// LoginViewModel struct
type LoginViewModel struct{
	BaseViewModel
	Errs []string
}

func (v *LoginViewModel) AddError (errs ...string) {
	v.Errs = append(v.Errs, errs...)
}

// LoginViewModelOp struct{}
type LoginViewModelOp struct{}

// GetVM func
func (LoginViewModelOp) GetVM() LoginViewModel {
	v := LoginViewModel{}
	v.SetTitle("Login")
	return v
}

func CheckLogin(username, password string) bool {
	user, err := model.GetUserByUsername(username)
	if err != nil {
		log.Println("Can not find username: ", username)
		log.Println("Here is error: ", err)
		return false
	}
	return user.CheckPassword(password)
}
