package main

import "fmt"

//Person 结构体字段如果是指针，map，slice，则默认值是nil，即还没有分配空间，如果要使用，需要先make
type Person struct {
	Name  string
	Age   int
	score [3]float64
	p     *string
	m     map[string]string
	slice []int
}

func main() {
	var person Person
	if person.p == nil {
		fmt.Println("指针属性p是nil")
	}
	if person.m == nil {
		fmt.Println("map类型m是nil")
	}
	if person.slice == nil {
		fmt.Println("slice类型slice是nil")
	}
	fmt.Println("person的值=", person)
	fmt.Println("---------------------------")

	var p1 = person
	//注意要是map字段，一定要先make
	p1.m = make(map[string]string)
	p1.m["key1"] = "value1" //修改p1的值，不影响person的值，这两个变量是相互独立的
	fmt.Println("person的值=", person)
	fmt.Println("p1的值=", p1)

	//p2是指针
	var p2 = new(Person)
	(*p2).Name = "tom" //通常写法 结构体指针.字段名
	p2.Name = "mary"   //go编译器作了转化，p2.Name等价于(*p2).Name
	fmt.Println(*p2)

	//p3是指针
	var p3 = &Person{}
	(*p3).Age = 30
	p3.Age = 20
	fmt.Println(*p3)

}
