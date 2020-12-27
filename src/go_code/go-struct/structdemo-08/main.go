package main

import "fmt"

//Circle struct
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) String() string {
	str := fmt.Sprintf("radius=[%v]", c.radius)
	return str
}

func main() {

	var c Circle = Circle{3.0}
	res := c.area()
	fmt.Printf("c的面积 = %.2f\n", res)
	//如果Cirlce实现了String()方法，使用fmt.Println输出实例时会自动调用变量实现的String()方法
	fmt.Println(c)
}
