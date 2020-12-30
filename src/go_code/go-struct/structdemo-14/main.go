package main
import (
	"fmt"
)

type A struct{
	Name string
}

type B struct{
	Name string
}

type C struct{
	A
	B
	// Name string
}

func (a A) test(){
	fmt.Println("A...")
}

func (b B) test(){
	fmt.Println("B...")
}

// func (c C) test(){
// 	fmt.Println("C...")
// }

func main(){
	var c = C{A{Name: "aa",},B{Name: "bb",},}
	// var c = C{A{Name: "aa",},B{Name: "bb",},"cc"}

	// fmt.Println(c.Name)	//如果这样写，会出错，因为编译器不知道Name是哪个结构体中的, 此时就必须指定是哪个匿名结构体名下的Name，应该像下面这么写
	fmt.Println(c.A.Name, c.B.Name)
	// 如果C本身就有自己的Name属性，就会调用自己的；要想调用A或B的，就需要指定名称了
	// fmt.Println(c.Name)
	
	// c.test()	// 因为A,B都有test()方法，而C本身没有，所以此时调用时编译器不知道调用谁的test方法会报错，必须指定调用谁的test方法，应该下面这么写
	c.A.test()
	c.B.test()
	// 如果C本身就有自己的test方法，就会调用自己的；要想调用A或B的，就需要指定名称了
	// c.test()
}