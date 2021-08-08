package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {

	var pChan chan Person = make(chan Person, 10)
	var num int
	for i := 0; i < 10; i++ {
		num = rand.Intn(30)
		pChan <- Person{Name: "name" + strconv.Itoa(num), Age: num, Address: "address" + strconv.Itoa(num)}
	}
	close(pChan)
	/*	1、channel的for-range遍历
		如果channel关闭了,就会正常遍历，遍历完后就会退出
		如果channel没有关闭，遍历就会报错deadlock
	*/
	for v := range pChan {
		fmt.Printf("%v %v %v\n", v.Name, v.Age, v.Address)
	}
	// 传统方法变量管道
	// var v Person
	// for i:= 0; i< 10; i++ {
	// 	v = <-pChan
	// 	fmt.Printf("%v %v %v\n", v.Name, v.Age, v.Address)
	// }
	fmt.Println("------------------------------")

	/*	2、channel的关闭
		channel可以通过close()关闭，channel关闭后就不能再写入数据了，即使channel未满也不能写入了
		channel关闭后仍然可以读取数据，数据读取完后再读取就会取到对应数据类型的默认值，如int channel就是0，string就是空串等

	*/
	var str chan int = make(chan int, 3)
	str <- 4
	// close(str)
	str <- 7
	str <- 3

	val := <-str
	close(str)
	val = <-str
	val = <-str
	val = <-str
	val = <-str
	val = <-str
	fmt.Println(val)

}
