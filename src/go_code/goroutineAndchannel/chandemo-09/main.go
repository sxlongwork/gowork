package main

import (
	"fmt"
)

// 形参为只写管道，防止误读
func writeData(intChan chan<- int){
	for i:= 0; i< 10; i++{
		intChan <- i
	}
	close(intChan)
}

// 形参为只读管道，防止误写
func readData(intChan <-chan int, exitChan chan<- bool){
	for {
		v, ok := <- intChan
		if !ok{
			break
		}
		fmt.Printf("%v\n", v)
	}
	exitChan <- true
}
func main(){
	// 默认定义的都为双向管道，
	var intChan chan int = make(chan int,3)
	var exitChan chan bool = make(chan bool, 3)

	go writeData(intChan)
	for i := 0; i< 2; i++{
		go readData(intChan, exitChan)
	}
	for i := 0; i< 2; i++{
		<- exitChan
	}
}

