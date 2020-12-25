package main

import "fmt"

func main() {
	var m = make(map[string]map[string]string)
	m["stu1"] = make(map[string]string)
	m["stu1"]["name"] = "tom"
	m["stu1"]["age"] = "20"
	stu2 := map[string]string{
		"name": "jack",
		"age":  "25",
	}
	m["stu2"] = stu2

	//遍历m
	for k, v := range m {
		fmt.Printf("%v\n", k)
		for k2, v2 := range v {
			fmt.Printf("\t%v=%v\n", k2, v2)
		}
	}

	//map长度
	fmt.Printf("map长度=%v", len(m))

}
