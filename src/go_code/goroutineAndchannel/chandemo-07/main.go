package main

import (
	"strconv"
	"fmt"
)

func writeData(numChan chan int){
	for i := 1; i<= 2000; i++{
		numChan <- i
		// fmt.Printf("写入数据 %v\n", i)
	}
	close(numChan)
}

func readData(numChan chan int, resChan chan string, exitChan chan int){
	for {
		v, ok := <- numChan
		if !ok {
			break
		}
		resChan <- ("res[" + strconv.Itoa(v) + "]=" + strconv.Itoa(getRes(v)))
		// fmt.Printf("读出数据 计算后写入\n")
	}
	n := <- exitChan
	exitChan <- (n + 1)
}

func getRes(n int) int{
	sum := 0
	for i := 1; i<= n; i++{
		sum += i
	}
	return sum
}

func main(){
	var numChan chan int = make(chan int, 200)
	var resChan chan string = make(chan string, 200)
	var exitChan chan int = make(chan int, 1)
	exitChan <- 0

	go writeData(numChan)
	for i := 1; i<= 8; i++ {
		go readData(numChan, resChan, exitChan)
	}
	go func(){
		for {
			v := <- exitChan
			if v == 8 {
				close(resChan)
				break
			}
			exitChan <- v
		}
	}()
	for v := range resChan{
		fmt.Printf("%v\n", v)
	}

}