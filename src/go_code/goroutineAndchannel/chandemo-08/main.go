package main

import (
	"fmt"
)

func writeData(intChan chan int){
	for i := 1; i<= 2000; i++{
		intChan <- i
		// fmt.Printf("写入数据 %v\n", i)
	}
	close(intChan)
}

func readData(intChan chan int, resChan chan int, exitChan chan bool){
	for {
		v, ok := <- intChan
		if !ok {
			break
		}
		if isPrime(v){
			resChan <- v
		}
		// fmt.Printf("写入素数 %v\n", v)
	}
	exitChan <- true
}

func isPrime(n int) bool{
	if n == 1{
		return false
	}
	for i := 2; i <= n/2; i++{
		if n % i == 0{
			return false
		}
	}
	return true
}

func main(){
	var intChan = make(chan int, 200)
	var resChan = make(chan int, 100)
	var exitChan = make(chan bool, 4)

	go writeData(intChan)
	for i := 1; i<= 4; i++{
		go readData(intChan, resChan, exitChan)
	}
	go func(){
		for i := 0; i < 4; i++{
			<- exitChan
		}
		close(resChan)
	}()
	for {
		v, ok := <- resChan
		if !ok{
			break
		}
		fmt.Printf("读取素数 %v\n", v)
	}

}