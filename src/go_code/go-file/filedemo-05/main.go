package main

import (
	"os"
	"fmt"
	"bufio"
)

func main(){
	//打开一个已存在的文件，将文件清空并写入数据
	filePath := "./test.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0640)
	if err != nil {
		fmt.Println("open file error, err =", err)
		return
	}

	//关闭文件
	defer file.Close()
	//创建writer，写入数据
	data := "你好 golang!\n"
	writer := bufio.NewWriter(file)
	for i :=0; i<10; i++{
		writer.WriteString(data)
	}
	// 写入数据是一定要执行Flush方法，将数据从缓冲区写入文件
	writer.Flush()
}