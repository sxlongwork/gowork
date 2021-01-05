package main

import (
	"go_code/go-project/projectdemo-03/userservice"
	"fmt"
)

type UserView struct{
	userservice userservice.UserService
	sel string
	flag bool
}

func (uv *UserView) List(){
	fmt.Println("------------------------客户列表------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	uv.userservice.List()
	fmt.Println("----------------------客户列表完成----------------------")
	fmt.Println()

}

func (uv *UserView) Add(){
	name := ""
	sex := ""
	age := 0
	mobile := ""
	email := ""
	fmt.Println("------------------------添加客户------------------------")
	fmt.Printf("姓名:")
	fmt.Scanln(&name)
	fmt.Printf("性别:")
	fmt.Scanln(&sex)
	fmt.Printf("年龄:")
	fmt.Scanln(&age)
	fmt.Printf("电话:")
	fmt.Scanln(&mobile)
	fmt.Printf("邮箱:")
	fmt.Scanln(&email)
	
	uv.userservice.Add(name,sex,age,mobile,email)
	fmt.Println("------------------------添加完成------------------------")
	fmt.Println()
}

func (uv *UserView) Delete(){
	id := 0
	sel := ""
	fmt.Println("------------------------删除客户------------------------")
	fmt.Printf("请选择待删除用户编号(-1退出):")
	fmt.Scanln(&id)
	if id == -1 {
		fmt.Println()
		return
	}
	fmt.Printf("确认是否删除(y/n):")
	fmt.Scanln(&sel)
	if sel == "y" {
		if  ok, index := uv.userservice.FindUserById(id); ok{
			uv.userservice.Delete(index)
			fmt.Println("------------------------删除完成------------------------")
		} else {
			fmt.Println("该客户信息不存在!")
			fmt.Println("------------------------删除失败------------------------")
		}
	}
	fmt.Println()
}

func (uv *UserView) Modify(){
	id := 0
	fmt.Println("------------------------修改客户------------------------")
	fmt.Printf("请选择待修改用户编号(-1退出):")
	fmt.Scanln(&id)
	if id == -1 {
		fmt.Println()
		return
	}
	if  ok, index := uv.userservice.FindUserById(id); ok{
		name := ""
		sex := ""
		age := 0
		mobile := ""
		email := ""
		fmt.Printf("姓名(%v):",uv.userservice.Users[index].Username)
		fmt.Scanln(&name)
		fmt.Printf("性别(%v):",uv.userservice.Users[index].Sex)
		fmt.Scanln(&sex)
		fmt.Printf("年龄(%v):",uv.userservice.Users[index].Age)
		fmt.Scanln(&age)
		fmt.Printf("电话(%v):",uv.userservice.Users[index].Mobile)
		fmt.Scanln(&mobile)
		fmt.Printf("邮箱(%v):",uv.userservice.Users[index].Email)
		fmt.Scanln(&email)

		uv.userservice.Modify(index,name, sex, age,mobile,email)
		fmt.Println("------------------------修改完成------------------------")
	} else {
		fmt.Println("该客户信息不存在!")
		fmt.Println("------------------------修改失败------------------------")
	}

}

func (uv *UserView) MainMenu(){
	for{
		fmt.Println("----------客户信息管理软件----------")
		fmt.Println("          1 添加客户")
		fmt.Println("          2 删除客户")
		fmt.Println("          3 修改客户")
		fmt.Println("          4 客户列表")
		fmt.Println("          5 退出")
		fmt.Println()
		fmt.Printf("请选择(1-5):")
		fmt.Scanln(&uv.sel)
		switch uv.sel {
		case "1":
			uv.Add()
		case "2":
			uv.Delete()
		case "3":
			uv.Modify()
		case "4":
			uv.List()
		case "5":
			uv.flag = true
		default:
			fmt.Println("输入有误, 请重新选择!")
		}
		if uv.flag{
			break
		}
	}
}

func main(){
	us := userservice.NewUserService()
	uv := UserView{us,"",false}
	uv.MainMenu()

}