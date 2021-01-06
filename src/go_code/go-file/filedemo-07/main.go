package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

func main(){
	filePath := "./test.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0640)
	if err != nil{
		fmt.Println("open file err, err =", err)
		return
	}

	defer file.Close()

	// 逐行读取数据并显示到终端
	reader := bufio.NewReader(file)
	for {
		data, err := reader.ReadString('\n')
		fmt.Print(data)
		if err == io.EOF{
			break
		}
	}
	//追加数据到文件中
	content := "hello 北京\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++{
		writer.WriteString(content)
	}
	writer.Flush()
}