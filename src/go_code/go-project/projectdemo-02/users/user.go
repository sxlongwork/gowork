package users

import (
	"fmt"
	"go_code/go-project/projectdemo-02/account"
)

type Users struct{
	useraccounts map[string]account.FamilyAccount
	sel string
	flag bool
	username string
}

func Start(user *Users){
	for {
		fmt.Println("----------欢迎来到家庭收支记账软件----------")
		fmt.Println("          1 注册")
		fmt.Println("          2 登录")
		fmt.Println("          3 退出")
		fmt.Println()
		fmt.Printf("          请选择(1-3):")
		fmt.Scanln(&user.sel)
		if user.sel == "1" {
			for {
				fmt.Printf("请输入用户名:")
				fmt.Scanln(&user.username)
				if _, ok := user.useraccounts[user.username]; ok{
					fmt.Println("该用户名已注册,请重新输入!")
				} else {
					user.useraccounts[user.username] = *account.GetFamilyAccount()
					break
				}
				fmt.Printf("请输入密码:")
				// fmt.Scanln(user.useraccounts[user.username].)
			}
			
		} else if user.sel == "2"{
			fmt.Println("登录")
		} else if user.sel == "3"{
			break
		} else {
			fmt.Println("选择有误, 请重新选择!")
		}
	}
	
}