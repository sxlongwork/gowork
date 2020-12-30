package main
import "fmt"

//全局匿名函数
var (
	fun1 = func (n1 int, n2 int) int {
		return n1 * n2
	}
)

func main(){

	//匿名函数使用方式1，在定义的时候调用，这种方式匿名函数只能调用一次
	var sum = func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("sum = ", sum)

	//匿名函数使用方式2，将匿名函数赋给变量，可以通过该变量调用
	var sub = func(n1 int, n2 int) int {
		return n1 - n2
	}

	fmt.Println(sub(20, 10))
	fmt.Println(sub(100,20))

	fmt.Println(fun1(2,8))
}