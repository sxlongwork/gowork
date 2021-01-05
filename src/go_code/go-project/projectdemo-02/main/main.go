package main

import (
	_ "go_code/go-project/projectdemo-02/account"
	"go_code/go-project/projectdemo-02/users"
)

func main(){
	// account.GetFamilyAccount().MainMenu()
	users.GetInstance().Start()
}