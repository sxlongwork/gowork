package main

import "fmt"

type Student struct {
	Name string
	Age  byte
}

func main() {
	var m = make(map[string]Student)
	stu1 := Student{"xiaoming", 20}
	m["01"] = stu1

	stu2 := new(Student)
	stu2.Name = "tom"
	stu2.Age = 23
	m["02"] = *stu2

	//遍历
	for _, v := range m {
		fmt.Println(v.Name, v.Age)
	}

}
