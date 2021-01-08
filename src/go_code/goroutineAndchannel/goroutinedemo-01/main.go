package main

import (
	"fmt"
	"time"
)

func test(){
	for i := 0; i < 10; i++ {
		fmt.Println("[test] hello world", i)
		time.Sleep(time.Second)
	}
}

func main(){
	// 开启一个goroutine
	go test()
	// 主线程
	for i := 0; i < 10; i++ {
		fmt.Println("[main] hello goland", i)
		time.Sleep(time.Second)
	}
}