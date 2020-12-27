package main

import "fmt"

type student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

//Visitor struct
type Visitor struct {
	name string
	age  byte
}

func (v Visitor) getPrice() {
	if v.age < 5 || v.age > 80 {
		fmt.Println("年龄不合适，为了安全，请不要玩!")
		return
	}
	if v.age > 18 {
		fmt.Printf("%v的年龄为:%v,门票价格为:20元\n", v.name, v.age)
	} else {
		fmt.Printf("%v的年龄为:%v,门票免费\n", v.name, v.age)
	}
}

func (s student) say() string {
	str := fmt.Sprintf("name=[%v],gender=[%v],age=[%v],id=[%v],score=[%v]",
		s.name, s.gender, s.age, s.id, s.score)
	return str
}
func main() {
	var stu = student{
		name:   "tom",
		gender: "male",
		age:    20,
		id:     1,
		score:  87.5,
	}
	res := stu.say()
	fmt.Println(res)

	var vis Visitor
	var name string
	var age byte
	for {
		fmt.Println("请输入姓名:")
		fmt.Scanln(&name)
		if name == "n" {
			fmt.Println("退出程序")
			break
		}
		vis.name = name
		fmt.Println("请输入年龄:")
		fmt.Scanln(&age)
		vis.age = age

		vis.getPrice()

	}
}
