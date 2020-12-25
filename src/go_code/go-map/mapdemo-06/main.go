package main

import "fmt"

func main() {

	var slice = make([]map[string]string, 1)

	el := make(map[string]string)
	el["name"] = "tom"
	el["age"] = "20"
	slice[0] = el
	fmt.Println(slice)

	//动态添加map的个数
	e2 := map[string]string{
		"name": "jack",
		"age":  "23",
	}
	slice = append(slice, e2) //追加到slice末尾
	fmt.Println(slice)
}
