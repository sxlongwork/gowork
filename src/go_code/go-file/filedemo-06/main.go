package main

import (
	"os"
	"fmt"
	"bufio"
)

func main(){
	filePath := "./test.log"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil{
		fmt.Println("open file err, err =", err)
		return
	}

	defer file.Close()

	data := "ABC ENGLISH!\n"
	writer := bufio.NewWriter(file)

	writer.WriteString(data)

	writer.Flush()
}