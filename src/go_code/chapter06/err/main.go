package main
import (
	"fmt"
	"errors"
)

func test(){
	//使用defer + recover来处理异常
	defer func(){
		err := recover()	//recover()为内建函数，可以捕获到异常
		if err != nil{		//说明捕获到异常
			fmt.Println("err = ", err)
		} else {
			fmt.Println("代码正常")
		}
	}()
	n1 := 2
	n2 := 0
	res := n1 / n2
	fmt.Println("res=", res)
}

func readConf(name string) (err error) {
	if name == "log.txt" {
		return nil
	} else {
		//自定义错误
		err := errors.New("文件打开错误，文件名称不对...")
		return err
	}
}

func test2() {
	err := readConf("log.jpg")
	if err != nil {
		//抛出错误，退出程序
		panic(err)
	}
	fmt.Println("test2后面的代码")
}

func main(){

	// test()
	test2()
	fmt.Println("main()下面的代码")
}
