package main

import (
	"fmt"
)

func writeData(intChan chan int){
	for i := 1; i <= 50; i++ {
		intChan <- i * 2
		fmt.Printf("写入数据 %v\n", i*2)
	}
	// 数据写完后关闭管道
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool){
	for {
		v, ok := <- intChan
		if !ok {
			break
		}
		fmt.Printf("读取数据 %v\n", v)
	}
	exitChan <- true
	close(exitChan)
}

func main(){
	var intChan chan int = make(chan int, 10)
	var exitChan chan bool = make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)
	
	for {
		_, ok := <- exitChan
		if !ok {
			break
		}
	}
	fmt.Println("main 退出")
}