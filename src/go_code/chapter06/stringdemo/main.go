package main

import (
	"fmt"
	"go_code/chapter06/stringdemo/arraylist"
	"go_code/chapter06/stringdemo/stack"
)

func main() {

	var list arraylist.List = new(arraylist.ArrayList)
	list.Append("aa")
	list.Append("bb")
	list.Append(12)
	list.Delete(1)
	// list.Clear()
	for i:=0; i<20 ;i++{
		list.Insert(1, "x3")
	}
	fmt.Println(list)

	str := make([]int, 10, 10)
	str[0] = 1
	str[1] = 3
	str = append(str, 10)
	fmt.Println(str)
	fmt.Println("--------------------------------------------")

	mystack := stack.NewStack()
	// mystack.Push(1)
	// mystack.Push(2)
	for i :=0; i<50;i++ {
		mystack.Push(i)
	}
	fmt.Println(mystack)
	fmt.Println(mystack.Size())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack)
	fmt.Println(mystack.Size())
}
