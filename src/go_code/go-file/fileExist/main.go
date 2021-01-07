package main

import (
	"os"
	"fmt"
)

func isExist(filePath string) (bool, error){
	_, err := os.Stat(filePath)
	if err == nil{
		return true, nil
	}
	if os.IsNotExist(err){
		return false, nil
	}
	return false, err
}

func main(){
	filePath := "./test.txt"
	if exist, _ := isExist(filePath); exist{
		fmt.Println("文件存在")
	} else {
		fmt.Println("文件不存在")
	}
	
}