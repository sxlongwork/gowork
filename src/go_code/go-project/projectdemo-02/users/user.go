package users

import (
	"fmt"
	"go_code/go-project/projectdemo-02/account"
)

type Users struct{
	useraccounts map[string]*account.FamilyAccount
	sel string
	// flag bool
}

func GetInstance() *Users{
	m := make(map[string]*account.FamilyAccount)
	return &Users{
		useraccounts: m,
		sel: "",

	}
}

func (user *Users) Start(){
	for {
		fmt.Println("----------欢迎来到家庭收支记账软件----------")
		fmt.Println("          1 注册")
		fmt.Println("          2 登录")
		fmt.Println("          3 退出")
		fmt.Println()
		fmt.Printf("          请选择(1-3):")
		fmt.Scanln(&user.sel)
		if user.sel == "1" {
			user.register()
		} else if user.sel == "2"{
			user.login()
		} else if user.sel == "3"{
			break
		} else {
			fmt.Println("选择有误, 请重新选择!")
		}
	}
}

func (user *Users) login(){
	username := ""
	pwd := ""
	fmt.Printf("用户名:")
	fmt.Scanln(&username)
	fmt.Printf("密 码:")
	fmt.Scanln(&pwd)
	fmt.Println(user.useraccounts)
	if v, ok := user.useraccounts[username]; ok && v.GetPwd() == pwd{
			fmt.Println("登录成功!")
			v.MainMenu()
	} else{
		fmt.Println("用户名或密码错误")
	}

}

func (user *Users) register(){
	username := ""
	pwd := ""
	for {
		fmt.Printf("请输入用户名:")
		fmt.Scanln(&username)
		if _, ok := user.useraccounts[username]; ok{
			fmt.Println("该用户名已注册,请重新输入!")
			continue
		} else {
			user.useraccounts[username] = account.GetFamilyAccount()
		}
		fmt.Printf("请输入密码:")
		fmt.Scanln(&pwd)
		if v, ok := user.useraccounts[username]; ok{
			v.SetPwd(pwd)
			fmt.Println(v)
			fmt.Println("注册成功, 请登录!")
			break
		}
	}
}