package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	// 定义源文件与目的文件路径
	srcFile := "./src.txt"
	dstFile := "./dst.txt"
	// 读取源文件内容
	data, err := ioutil.ReadFile(srcFile)
	if err != nil{
		fmt.Println("open filee error, err =", err)
		return
	}
	// WriteFile: 函数向filename指定的文件中写入数据。如果文件不存在将按给出的权限创建文件，否则在写入数据之前清空文件。
	// 将从源文件读取的数据data写入目的文件
	err = ioutil.WriteFile(dstFile, data, 0640)
	if err != nil{
		fmt.Println("write file err, err =", err)
	}

}