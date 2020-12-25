package main

import "fmt"

func main() {
	//使用map存储3个学生信息，每个学生信息包括姓名和性别
	//可以使用map[string]map[string][string]
	var m = make(map[string]map[string]string, 3)

	stu1 := make(map[string]string)
	stu1["name"] = "xiaoming"
	stu1["sex"] = "m"
	m["stu1"] = stu1

	stu2 := map[string]string{
		"name": "tom",
		"sex":  "m",
	}
	m["stu2"] = stu2

	m["stu3"] = make(map[string]string)
	m["stu3"]["name"] = "jack"
	m["stu3"]["sex"] = "f"

	//遍历输出
	for k1, v1 := range m {
		fmt.Println(k1)
		for k2, v2 := range v1 {
			fmt.Printf("\tk=%v, v=%v\n", k2, v2)
		}
		fmt.Println()
	}
}
