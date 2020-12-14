package main

import "fmt"

func main() {
	//map的初始化及使用
	// 1、要求所有的key的数据类型相同，所有value数据类型相同(注：key与value可以有不同的数据类型)
	// 2、map中的每个key在keys的集合中是唯一的，而且需要支持 == or != 操作
	// 3、key的常用类型：int, rune, string, 结构体(每个元素需要支持 == or != 操作), 指针, 基于这些类型自定义的类型

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

	//查找键值是否存在,若key在map中，返回对应的值;若key不在map中，返回value类型的默认值
	if v, ok := m3["brith"]; ok {
		fmt.Println(ok)
		fmt.Println(v)
	} else {
		fmt.Println(ok)
		fmt.Println(v)
		fmt.Println("key not exist")
	}
	//遍历map
	for k, v := range m3 {
		fmt.Printf("key=%v, value=%v\n", k, v)
	}

	//map的增删查改
	// 添加 修改,key不在map中，就是添加，在map中就是修改
	m1["nick"] = "www"
	m1["name"] = "xiaoqian"
	//查, key在map中，返回值给v，ok=true; key不在map中，返回value类型的默认值给v，ok=false
	v, ok := m1["age"]
	fmt.Printf("v=%v, ok=%v", v, ok)
	//删除,delete(map, key),key在map中，删除key及对应的value；key不在map中，则不执行任何操作
	delete(m1, "nick")
}
