package main

import "fmt"

func main() {

	var m map[string]string
	m = make(map[string]string)
	m["name"] = "xiaoming" //新增元素
	m["sex"] = "male"
	fmt.Println("新增元素后", m)
	m["name"] = "tom" //修改元素
	fmt.Println("修改name后", m)

	//查找元素
	v, ok := m["name"]
	if ok {
		fmt.Println("name在map中,value为", v)
	} else {
		fmt.Println("name不在map中")
	}

	//删除元素
	delete(m, "name")
	fmt.Println("删除name后", m)

}
