package model

import "fmt"

/*
要求：账号必须6-20位；密码必须6位；余额必须大于20
*/
type account struct {
	name string
	pwd  string
	bal  float64
}

/*
GetAccount 对外提供获取实例函数
*/
func GetAccount(name string, pwd string, bal float64) *account {
	if len(name) < 6 || len(name) > 10 {
		fmt.Println("账号位数不对")
		return nil
	}
	if len(pwd) != 6 {
		fmt.Println("密码位数不对")
		return nil
	}
	if bal <= 20 {
		fmt.Println("")
		return nil
	}
	return &account{
		name: name,
		pwd:  pwd,
		bal:  bal,
	}
}

func (acc *account) SetName(name string) {
	acc.name = name
}

func (acc *account) GetName() string {
	return acc.name
}

func (acc *account) SetPwd(pwd string) {
	acc.pwd = pwd
}

func (acc *account) GetPwd() string {
	return acc.pwd
}

func (acc *account) SetBal(bal float64) {
	acc.bal = bal
}

func (acc *account) GetBal() float64 {
	return acc.bal
}
