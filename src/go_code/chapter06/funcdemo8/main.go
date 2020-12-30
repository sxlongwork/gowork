package main
import "fmt"

func sum(n1,n2 float32) float32{
	return n1+n2
}

//交换两个变量的值
func swap(n1 *int, n2 *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
}


func main(){

	res := sum(1,2)
	fmt.Println(res)

	a := 10
	b := 20
	swap(&a, &b)
	fmt.Println("a=", a)
	fmt.Println("b=", b)
}