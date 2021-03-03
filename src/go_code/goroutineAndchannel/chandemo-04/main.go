package main

import (
	"math/rand"
	"strconv"
	"fmt"
)

type Person struct{
	Name string
	Age int
	Address string
}

func main(){

	var pChan chan Person = make(chan Person, 10)
	var num int
	for i := 0; i< 10; i++ {
		num = rand.Intn(30)
		pChan<- Person{Name: "name"+strconv.Itoa(num),Age: num, Address: "address"+strconv.Itoa(num)}
	}
	// close(pChan)
	// // for range变量管道时，如果管道没有close，会报错
	// for v := range pChan{
	// 	fmt.Printf("%v %v %v\n", v.Name, v.Age, v.Address)
	// }
	// 传统方法变量管道
	var v Person
	for i:= 0; i< 10; i++ {
		v = <-pChan
		fmt.Printf("%v %v %v\n", v.Name, v.Age, v.Address)
	}
}