package main
import "fmt"

//函数在执行过程中，遇到defer时，暂时不执行，而是把defer后面的语句压到独立的栈中，待函数执行完毕后，再从栈中遵从先进后出的方式执行压入的语句
func sum(n1 int, n2 int) int{
	defer fmt.Println("n1=", n1)	//第三执行
	defer fmt.Println("n2=", n2)	//第二执行
	sum := n1 + n2
	fmt.Println("sum=", sum)	//第一执行
	return sum
}

func test(){
	//打开文件
	file := openFile()
	defer file.close()	//关闭文件
	// 执行其他代码
}

func dbConnect(){
	//打开数据库连接
	connect := openDb()
	defer connect.close()	//关闭数据库连接
	// 执行其他代码
}

func main(){
	fmt.Println()

	res := sum(10, 20)
	fmt.Println("res=", res)	//第四执行
}