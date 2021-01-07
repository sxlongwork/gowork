package main

import (
	"fmt"
	"flag"
)

func main(){
	// 注意返回的是对应类型的指针,这种函数底层调用的也是对应的XxxVar函数
	var name = flag.String("name", "tom", "user name")
	var age = flag.Int("age", 10, "user age")
	// 也可以使用以下函数
	var mobile string
	flag.StringVar(&mobile, "mobile", "13200001111", "user mobile")

	// 该方法从os.Args[1:]中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。
	// 未注册却使用flag -help时，会返回ErrHelp。
	flag.Parse()

	fmt.Println(*name, *age, mobile)
}