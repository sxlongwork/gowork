package main

import "fmt"

type B interface{
	testB()
}

type C interface{
	testC()
}

type A interface {
	B
	C
	// testA()
}

type D interface{

}

type Inte int

func (i Inte) testB(){

}
func (i Inte) testC(){

}


func main(){
	var a Inte = 3

	var ain A = a
	fmt.Printf("%p, %v\n", &a, a)
	fmt.Printf("%p, %v\n", &ain, ain)

	// var d D = 3
	// var d D = "hello"
	var d D = [2]int{2,4}
	fmt.Println(d)

}