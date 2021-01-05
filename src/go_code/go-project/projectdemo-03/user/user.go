package user

import (
	"fmt"
)

type User struct{
	Id int
	Username string
	Sex string
	Age int
	Mobile string
	Email string
}


func NewUser(id int,username string,sex string,age int,mobile string,email string) User{
	return User{
		Id:id,
		Username: username,
		Sex:sex,
		Age: age,
		Mobile: mobile,
		Email: email,
	}
}

func (user *User) List(){
	str := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v",user.Id,user.Username,user.Sex,user.Age,user.Mobile,user.Email)
	fmt.Println(str)
}