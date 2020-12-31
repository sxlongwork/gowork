package main

import (
	"fmt"
)

func main(){

	//定义变量记录用户的选择
	var sel int
	//定义变量记录是否退出
	var flag = false
	//当前余额
	var bal = 1000.0
	//支出/收入金额
	var money float64
	//收入/支出说明
	var note string
	//定义变量记录收支明细
	var details string = "收支\t余额\t金额\t说明"
	//定义变量记录使用有收支记录
	var record = false
	for {
		fmt.Println("----------家庭收支记账软件----------")
		fmt.Println("          1 收支明细")
		fmt.Println("          2 登记收入")
		fmt.Println("          3 登记支出")
		fmt.Println("          4 退出")
		fmt.Println()
		fmt.Printf("          请选择(1-4):")
		fmt.Scanln(&sel)
		switch sel{
		case 1:
			if record{
				fmt.Println(details)
			} else {
				fmt.Println("您还没有收支记录，来记一笔吧!")
			}
			fmt.Println()
		case 2:
			fmt.Printf("本次收入金额:")
			fmt.Scanln(&money)
			fmt.Printf("本次收入说明:")
			fmt.Scanln(&note)

			bal += money
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", bal, money, note)
			record = true
		case 3:
			fmt.Printf("本次支出金额:")
			fmt.Scanln(&money)
			if money > bal {
				fmt.Println("余额不足,无法支出!")
				break
			}
			fmt.Printf("本次支出说明:")
			fmt.Scanln(&note)

			bal -= money
			details += fmt.Sprintf("\n支出\t%v\t%v\t%v", bal, money, note)
			record = true
		case 4:
			var key string
			for {
				fmt.Println("您确定退出吗? (y/n):")
				fmt.Scanln(&key)
				if key == "y"{
					flag = true
					break
				} else if key == "n" {
					break
				} else{
					fmt.Println("您的输入有误，请重新输入")
				}
			}
		default:
			fmt.Println("您的选择有误，请重新选择!")
		}
		if flag {
			break
		}

	}
}