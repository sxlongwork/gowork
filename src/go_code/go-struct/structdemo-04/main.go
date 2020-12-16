package main

import "fmt"

type human struct {
	name string
}

//结构体之间的"继承", 嵌入结构体
type student struct {
	age   byte
	score float64
	human
}

//判断[5]int数组中是否含有重复元素，有就返回true; 没有就返回false
func containsDuplicate(arr [5]int) bool {
	var m = make(map[int]struct{})
	for _, v := range arr {
		if _, ok := m[v]; ok == true {
			return true
		}
		m[v] = struct{}{}
	}
	return false
}

func main() {
	var s student
	// fmt.Println("s =", s)	//s = {0 0 {}}
	s.age = 23
	s.score = 90.5
	// s.human.name = "xiaoming"
	s.name = "xiaoming"   //嵌入结构体相当于把字段直接拷贝过来了，操作字段时无需写成s.human.name
	fmt.Println("s =", s) //s = {23 90.5 {xiaoming}}
	s.name = "tom"
	fmt.Println("s =", s) //s = {23 90.5 {tom}}

	var m = make(map[string]struct{})
	m["k1"] = struct{}{}
	fmt.Println(m)

	//struct的特殊用法，用于判断数组或切片中是否含有重复元素
	//map[int]struct{},由于struct{}地址为空，不关心内容，这样map值就不会占用内存，map就变成了了set
	var arr = [5]int{2, 3, 2, 5, 6}
	if flag := containsDuplicate(arr); flag == true {
		fmt.Println("有重复元素")
	} else {
		fmt.Println("没有重复元素")
	}
	fmt.Println()
}
