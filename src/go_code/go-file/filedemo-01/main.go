package main
import (
	"fmt"
	"os"
)

func main(){
	//打开一个文件，使用os.Open(name string)函数,返回*File，error;使用该方法打开的文件具有O_RDONLY模式
	file, error := os.Open("./test.txt")

	//函数执行完后关闭文件
	defer file.Close()
	if error != nil{
		fmt.Println("open file failed, err=", error)
		return
	}
	fmt.Println(*file)
}