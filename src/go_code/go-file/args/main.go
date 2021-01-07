package main

import (
	"os"
	"fmt"
)

func main(){
	// 1、os.Args是一个字符串切片
	// 2、存储所有的命令行参数，其中下标为0的位置存储执行文件的路径
	// 3、os.Args[1:]存储用户传入的参数
	for index, _ := range os.Args{
		fmt.Printf("os.Args[%v]=%v\n", index, os.Args[index])
	}
	
}