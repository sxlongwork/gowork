package main

import (
	"fmt"
	"runtime"
)

func main(){
	// 获取当前系统CPU的逻辑数
	num := runtime.NumCPU()
	fmt.Println("CPU数目 =", num)
	// 设置几个cpu运行, 这里设置num-2个cpu运行程序
	runtime.GOMAXPROCS(num - 2)
	fmt.Println("all ok")
}