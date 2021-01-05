package userservice

import (
	"go_code/go-project/projectdemo-03/user"
)

type UserService struct{
	Users []user.User
	Id int
}

func NewUserService() UserService{
	user := user.NewUser(1,"tom","male",21,"132","tom@qq.com")
	userService := UserService{}
	userService.Users = append(userService.Users,user)
	userService.Id = 1
	return userService
}

func (us *UserService) List(){
	for i:= 0;i < len(us.Users); i++{
		us.Users[i].List()
	}
}

func (us *UserService) Add(name string,sex string,age int,mobile string,email string) {
	us.Id++
	user := user.NewUser(us.Id,name,sex,age,mobile,email)
	us.Users = append(us.Users, user)
}

func (us *UserService) Delete(index int){
	us.Users = append(us.Users[:index],us.Users[index+1:]...)
}

func (us *UserService) FindUserById(id int) (bool,int){
	for i:= 0;i < len(us.Users); i++{
		if us.Users[i].Id == id{
			return true, i
		}
	}
	return false, -1
}

func (us *UserService) Modify(index int,name string,sex string,age int,mobile string,email string){
	if name != ""{
		us.Users[index].Username = name
	}
	if sex != ""{
		us.Users[index].Sex = sex
	}
	if age != 0{
		us.Users[index].Age = age
	}
	if mobile != ""{
		us.Users[index].Mobile = mobile
	}
	if email != ""{
		us.Users[index].Email = email
	}
}