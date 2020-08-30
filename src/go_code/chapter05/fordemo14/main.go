package main

import "fmt"

func main() {
	//打印1-100所有是9的倍数的整数的个数及总和

	count := 0
	sum := 0
	for i := 1; i <= 100; i++ {
		if i % 9 == 0 {
			// fmt.Println("i =", i)
			count++
			sum += i
		}
	}
	fmt.Println("count =", count, "sum =", sum)

	var n int = 10
	for i := 0; i <= n; i++ {
		fmt.Printf("%d + %d = %d\n", i, n - i, n)
	}
}