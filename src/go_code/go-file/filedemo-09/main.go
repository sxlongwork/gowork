package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"unicode"
)

type Data struct{
	ch int
	num int
	space int
	others int
	han int
}

func main(){
	var data Data

	filePath := "./test.txt"

	//打开文件
	file, err := os.Open(filePath)
	if err != nil{
		fmt.Println("open file error, err =", err)
	}
	// 关闭文件
	defer file.Close()
	// 逐行读取文件，并统计
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		sli := []rune(str)
		for _, value := range sli{
			if unicode.Is(unicode.Han, value){
				data.han++
			}
			switch {
			case value >=65 && value <= 90:
				fallthrough
			case value >= 95 && value <=122:
				data.ch++
			case value >=48 && value <= 57:
				data.num++
			case value == 32:
				data.space++
			default:
				data.others++
			}
		}
		if err == io.EOF{
			break
		}

	}
	fmt.Printf("字母：%v\n数字：%v\n空格：%v\n汉字：%v\n其他：%v\n", data.ch,data.num,data.space,data.han,data.others)
}