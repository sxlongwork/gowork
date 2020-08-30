package utils

import "fmt"

//为了让其他文件调用该函数，函数名称的首字母要大写，小写代表私有，只能在本文中使用
func Cal(n1 float64, n2 float64, oper byte) float64 {
	var res float64
	switch oper {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("操作符错误！")
	}
	return res
}