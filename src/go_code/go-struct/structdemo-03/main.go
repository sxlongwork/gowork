package main

import "fmt"

//结构体嵌套
type person struct {
	id    int
	human struct {
		sex    bool
		weight float64
	}
}

func main() {
	//匿名结构体
	s1 := struct {
		id   int
		name string
	}{
		id:   20,
		name: "tom",
	}
	fmt.Println("s1=", s1)

	s2 := person{}
	fmt.Println(s2)
	s2.id = 1
	s2.human.sex = true
	s2.human.weight = 80.2
	fmt.Println("s2=", s2.id, s2.human.sex, s2.human.weight)

	//结构体之间的赋值
	s3 := s2
	fmt.Println(s2 == s3) //true
	s3.id = 2
	s3.human.weight = 62
	fmt.Println("s3=", s3)
	fmt.Println("s2=", s2)
	fmt.Println(s2 == s3) //false
}
