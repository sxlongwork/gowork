package main

import (
	"io/ioutil"
	"fmt"
)

func main(){
	// 文件需已存在，不存在会报错
	filePath := "./test.txt"
	//data为[]byte类型，成功的调用返回的err为nil而非EOF
	//不涉及文件的显式打开，所以也没有文件的显式关闭
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read file err, err =", err)
		return
	}
	fmt.Print(string(data))
}