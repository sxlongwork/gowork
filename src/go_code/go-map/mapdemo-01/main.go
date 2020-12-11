package main

import "fmt"

func main() {
	//map的初始化及使用
	//1.先声明，再make创建，最后赋值
	var m1 map[string]string
	m1 = make(map[string]string)
	m1["name"] = "xiaoming"
	m1["age"] = "20"
	//2.直接make创建，再赋值
	m2 := make(map[string]int)
	m2["k1"] = 2
	m2["k2"] = -4
	//3.初始化+赋值,不需要make
	m3 := map[string]bool{
		"info":  true,
		"error": false,
	}

	//查找键值是否存在
	if v, ok := m1["brith"]; ok {
		fmt.Println(ok)
		fmt.Println(v)
	} else {
		fmt.Println(ok)
		fmt.Println("key not exist")
	}
	//遍历map
	for k, v := range m3 {
		fmt.Printf("key=%v, value=%v\n", k, v)
	}
}
