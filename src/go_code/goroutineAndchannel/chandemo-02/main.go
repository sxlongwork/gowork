package main

import (
	"fmt"
)

func main(){
	//1. 定义及初始化
	// 定义了该chan的容量是2，最多只能放两个指定类型的值
	var intchan chan int
	intchan = make(chan int, 2)	

	//2. intchan的值是一个地址，它本身也有一个地址，这两个值不同
	fmt.Printf("intchan的值 = %v, intchan的地址是 = %v\n", intchan, &intchan)

	//3. 向管道中存放值
	intchan <- 20
	var num int = 100
	intchan <- num
	fmt.Printf("存放数据后, 此时 len(intchan) = %v, cap(intchan) = %v\n", len(intchan), cap(intchan))

	// 向管道中存放值的数量不能超过定义时的容量，该容量不想map一样会自动增大
	// intchan <- 0	// fatal error: all goroutines are asleep - deadlock!

	//4. 从管道中读取中，如上面存入了20,100，现在要获得100这个值
	<- intchan				// 将上面存放的第一个数20丢掉
	value := <- intchan		// 获取到第二个值100，此时读取完后该管道为空了
	fmt.Printf("从管道中获得value = %v\n", value)
	fmt.Printf("读取数据后, 此时 len(intchan) = %v, cap(intchan) = %v\n", len(intchan), cap(intchan))

	//5. 在没有使用协程的情况下, 如果已经知道管道中的数据已经全部取出，再取就会报错deadlock
	// <- intchan	// fatal error: all goroutines are asleep - deadlock!

}