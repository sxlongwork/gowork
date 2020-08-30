package main

import (
	"fmt"
)
func main() {
	//goto使用

	fmt.Println("xianqian--01")
	fmt.Println("xianqian--02")
	goto here
	fmt.Println("xianqian--03")
	fmt.Println("xianqian--04")
	fmt.Println("xianqian--05")
	here:
	fmt.Println("xianqian--06")
	fmt.Println("xianqian--07")
}