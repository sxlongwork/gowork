package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

func main(){

	//打开一个已存在的文件test.txt
	filePath := "./test.txt"
	file, err := os.Open(filePath)

	defer file.Close()	//关闭文件

	if err != nil{
		fmt.Println("open file err, err =",err)
		return
	}
	//NewReader创建一个具有默认大小缓冲、从r读取的*Reader
	reader := bufio.NewReader(file)
	for {
		//每次读取遇到换行符'\n'停止,即逐行读取；返回一个包含已读取的数据和'\n'字节的字符串
		str, err := reader.ReadString('\n')
		fmt.Print(str)
		//读取到文件末尾即io.EOF时退出
		if err == io.EOF{
			break
		}
	}
}