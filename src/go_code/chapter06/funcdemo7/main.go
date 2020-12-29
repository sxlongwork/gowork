package main
import "fmt"

// func test(n1 int) int {
// 	n2 := n1
// 	return n2
// }

func test1(n1 *int) int{
	*n1 = *n1 + 2
	return *n1
}

type myfunc = func(*int) int 

func test2(funcv myfunc, n1 *int) int {
	return funcv(n1)
}

//可以对函数返回值命名
func test3(n1 int, n2 int) (sum int, sub int){
	sum = n1 + n2	//注意此处 = 不能写为:=
	sub = n1 - n2
	return		//此时return处就不需要指定返回的变量了
}


//Go支持可变参数
func sum(n1 int, args... int) int{
	sum := n1
	for i := 0; i <len(args); i++ {
		sum += args[i]
	}
	return sum
}

func main() {
	// res := test(3)
	// fmt.Println("res=", res)

	// n1 := 3
	// test1(&n1)
	// fmt.Println(n1)

	// 在Go中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了。通过该变量可以对函数进行调用
	// mytest := test1
	// n2 := 10
	// mytest(&n2)
	// fmt.Printf("mytest的数据类型是%T\n",mytest)
	// fmt.Println(n2)

	//函数既然是一种数据类型，因此在Go红，函数可以作为形参，并且调用

	n := 100
	res := test2(test1, &n)
	fmt.Println("res = ", res)

	//基本语法：type 自定义数据类型名 数据类型         自定义数据类型
	// type myfunc = func(*int) int 

	res1 := sum(10, 100, 10, 30)
	fmt.Println("res1= ", res1)

}