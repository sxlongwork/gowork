package main
import "fmt"

var name = getName()	//最先执行

func getName() string{
	fmt.Println("test...")
	return "xianqian"
}

func init(){		//第二执行
	fmt.Println("init...")
}

func main(){		//第三执行
	fmt.Println("main...name=", name)
}