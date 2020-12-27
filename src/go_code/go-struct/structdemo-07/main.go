package main

import "fmt"

// Person struct
type Person struct {
	Name string
	Age  byte
}

// Dog struct
// type Dog struct{}

func (p Person) speak() {
	fmt.Printf("%v是一个好人!!!\n", p.Name)
}

func (p Person) jisuan() {
	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += i
	}
	fmt.Println("1+2+...+1000=", sum)
}

func (p Person) jisuan2(n int) {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	fmt.Println("1+2+...+n=", sum)
}

func (p Person) getSum(n1, n2 float64) float64 {
	return n1 + n2
}

func main() {
	// var dog Dog
	var person Person
	person.Name = "tom"
	// speak()	//不能直接调用，必须是Person类型的实例才能调用
	// dog.speak()	//Dog类型的实例也无法调用
	person.speak()
	person.jisuan()
	person.jisuan2(10)
	res := person.getSum(1.0, 2.5)
	fmt.Println("getSum(1.0,2.5) =", res)
}
