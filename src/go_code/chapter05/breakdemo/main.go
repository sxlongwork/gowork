package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	//随机数的生成
	rand.Seed(time.Now().Unix())
	// num := rand.Intn(100) + 1	//生成[1-100]的随机数
	// fmt.Println(num)

	count := 0
	for {
		num := rand.Intn(100) + 1
		fmt.Println("num =", num)
		count++
		if num == 99 {
			fmt.Printf("共随机生成了%d次才为99.\n", count)
			break
		}
	}

	//100以内的数求和，求出当和第一次大于20的当前数
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
		if sum > 20 {
			fmt.Println("当和第一次大于20的当前数 =", i)
			break
		}
	}

	//三次登录机会，用户名密码成功提示登录成功; 失败提示还有几次机会

	total := 3
	user := ""
	password := ""
	for i := 1; i <= total; i++ {
		fmt.Println("请输入用户名：")
		fmt.Scanln(&user)
		fmt.Println("请输入密码：")
		fmt.Scanln(&password)
		if user == "xianqian" && password == "xianqian" {
			fmt.Println("登录成功！")
			break
		} else {
			fmt.Printf("你还有%d次机会", total - i)
		}
	}
}