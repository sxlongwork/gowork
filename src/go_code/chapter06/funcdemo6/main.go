package main

import (
	"fmt"
)

//猴子吃桃问题
//第10天，剩余1个，eat(10) = 1
//第9天，剩余eat(9) = (eat(10) + 1) *2
// ...
// eat(n) = (eat(n+1) + 1) * 2
//求第一天猴子有多少桃子
//n为第几天，当n < 1或n > 10是，没有桃子 

func eat(n int) int {
	if n < 1 || n > 10 {
		fmt.Println("此时猴子没有桃子了")
		return 0
	} 
	if n == 10 {
		return 1
	} else {
		return (eat(n + 1) + 1) *2
	}
}

func main(){
	res := eat(1)
	fmt.Println("res =", res)
}