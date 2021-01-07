package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

func copy(dstFile string, srcFile string) (written int64, err error){

	srcfile, err := os.Open(srcFile)
	if err != nil{
		fmt.Println("open file error, err =", err)
		return
	}
	defer srcfile.Close()
	// 创建reader
	reader := bufio.NewReader(srcfile)

	dstfile, err := os.OpenFile(dstFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil{
		fmt.Println("open file error, err =", err)
		return
	}
	
	// 创建writer
	writer := bufio.NewWriter(dstfile)
	defer dstfile.Close()

	return io.Copy(writer, reader)
}

func main(){
	f1 := "D:/timg.jpeg"
	f2 := "E:/timg.jpeg"
	copy(f2, f1)
}