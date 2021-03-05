package main

import (
	"fmt"
	"time"
)

func writeData(intChan chan<- int){
	defer func(){
		err := recover()
		if err != nil{
			fmt.Println(err)
		}
	}()
	for i := 1; i< 15; i++{
		intChan <- i
		fmt.Printf("写入 %v\n", i)
	}
	// close(intChan)
}

func main(){
	var intChan chan int = make(chan int,10)

	go writeData(intChan)
	// for {
	// 	v, ok := <- intChan
	// 	if !ok{
	// 		break
	// 	}
	// 	fmt.Printf("读取 %v\n", v)
	// }
	label:
	for{
		time.Sleep(time.Second)
		select {
		case v := <- intChan:
			fmt.Printf("读取 %v\n", v)
		default:
			fmt.Printf("没数据\n")
			break label
		}
	}
	

}