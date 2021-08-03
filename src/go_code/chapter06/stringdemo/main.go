package main

import (
	"fmt"
	"go_code/chapter06/stringdemo/arraylist"
	"go_code/chapter06/stringdemo/stack"
	"go_code/chapter06/stringdemo/btree"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
)


type A struct{
	name string
}
/*
去除数组重复元素
输入：1 1 1 1 3 4 5 3
输出：1 3 4 5 
*/
func test() {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	datalist := strings.Fields(string(data))

	var numlist []int = make([]int, 0,10)
	for _, v := range datalist {
		num, _ := strconv.Atoi(v)
		numlist = append(numlist, num)
	}
	sort.Sort(sort.IntSlice(numlist))

	for i := 0; i< len(numlist)-1; i++ {
		if numlist[i] ^ numlist[i+1] == 0 {
			numlist = append(numlist[:i], numlist[i+1:]...)
			i--
		}
	}
	for _, v := range numlist{
		fmt.Print(v, " ")
	}
}

func main1() {

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
	fmt.Println("--------------------------------------------")

	var m map[string]string = make(map[string]string, 1)
	m["aa"] = "aa"
	fmt.Printf("%p\n",m)
	var a A = A{name:"oh"}
	fmt.Printf("%v",a)
	fmt.Println("--------------------------------------------")
}

func main(){

	var root  = btree.CreateBtree(5)
	root.Lchild = btree.CreateBtree(4)
	root.Rchild = btree.CreateBtree(6)
	root.Lchild.Lchild = btree.CreateBtree(3)
	root.Lchild.Rchild = btree.CreateBtree(4)

	node := root.FindNode(root, 4)
	if node != nil {
		fmt.Println(node.Data)
	}
	fmt.Println(root.GetHeight(root))

	reader := bufio.NewReader(os.Stdin)

	data,_,_ := reader.ReadLine()
	fmt.Printf("%T\n",data)
	fmt.Println(len(data))
}
