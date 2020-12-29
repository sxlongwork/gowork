package main
import (
	"fmt"
	"strings"
)
//

//该函数以文件后缀名为形参，返回一个闭包，该闭包接受一个文件名为参数，如果文件名含有该后缀，该返回文件名，如果没有该后缀则加上后缀返回
func makeSuffix(suffix string) func (string) string {

	return func(name string) string {
		if strings.HasSuffix(name, suffix) {
			return name
		} else {
			return name + suffix
		}
	}
	
}

func main(){
	fmt.Println("")
	var fun = makeSuffix(".jpg")
	fmt.Println(fun("test.jpg"))
	fmt.Println(fun("test"))
}