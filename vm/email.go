package vm

import (
	"blog/config"
	"blog/model"
)

type EmailViewModel struct {
	Username string
	Token 	 string
	Server	 string
}

type EmailViewModelOp struct {}

// GetVM func
func (EmailViewModelOp) GetVM(email string) EmailViewModel {
	v := EmailViewModel{}
	u, _ := model.GetUserByEmail(email)
	v.Username = u.Username
	toke, _ := u.GenerateToken()
	v.Token = toke
	v.Server = config.GetServerURL()
	return v
}
