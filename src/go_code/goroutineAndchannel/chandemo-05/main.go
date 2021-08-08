package main

import (
	"fmt"
	"time"
)

func main() {
	var intChan chan int = make(chan int, 3)
	intChan <- 2
	intChan <- 4
	close(intChan)
	// 虽然intChan容量为3，但是放入两个数据后关闭了管道，所以不能再向管道放入数据，否则报错
	// intChan <- 6 // panic: send on closed channel

	fmt.Println("ok")
	// 管道关闭后，读取数据是正常的
	n := <-intChan
	fmt.Println(n)

	var intChan2 chan int = make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan2 <- i * 2
	}
	close(intChan2)
	// 使用for-range遍历管道，遍历前一定要close管道
	for v := range intChan2 {
		fmt.Println(v)
	}
	fmt.Println("-----------------------------")
	var stringchan chan string = make(chan string, 3)
	go func() {
		stringchan <- "abc"
		stringchan <- "bcd"
		stringchan <- "cde"
		// close(stringchan)
	}()
	for {
		// if v, ok := <-stringchan; ok {
		// 	fmt.Println(v)
		// } else {
		// 	break
		// }
		time.Sleep(time.Second * 1)
		select {
		case v := <-stringchan:
			fmt.Println(v)
		default:
			fmt.Println("no data")
			return
		}

	}
}
