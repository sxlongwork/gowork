package main
import "fmt"

func myprint(n int){
	for i := 1; i <= n; i++ {
		for j := 1; j<= i; j++ {
			fmt.Printf("%d * %d = %d ", j, i, j*i)
		}
		fmt.Println()
	}
}

func main(){
	myprint(9)
}