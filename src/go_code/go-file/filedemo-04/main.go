package main

import (
	"os"
	"fmt"
	"bufio"
)

func main(){
	//打开一个文件，没有则创建，文件可写
	filePath := "./test.txt"
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil{
		fmt.Println("open file error, err =", err)
		return
	}
	//关闭文件
	defer file.Close()

	writer := bufio.NewWriter(file)
	data := "hello goloang\n"
	for i := 0; i<5; i++{
		writer.WriteString(data)
	}
	// 因为是带缓冲的写入，所以数据会先写入缓冲，必须执行Flush方法将数据写入文件；如果不执行，则文件中不会有数据
	writer.Flush()

}