package main

import (
	"fmt"
	"go_code/go-struct/structdemo-12/model"
)

func main() {

	var acc = model.GetAccount("xiaoming", "123456", 30.0)
	if acc != nil{
		fmt.Println(*acc)
	}
	acc.SetName("tom---")
	acc.SetPwd("666666")
	acc.SetBal(93.0)
	fmt.Println(acc.GetName(),acc.GetPwd(),acc.GetBal())
}
