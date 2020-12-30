package main
import "fmt"


func addUpper() func (int) int {
	var n = 10
	return func (a int) int {
		n = n + a
		return n
	}
}

func main(){

	fun := addUpper()

	fmt.Println(fun(1))	//11
	fmt.Println(fun(2))	//13
	fmt.Println(fun(3))	//16
}