package main

import "fmt"

// MethodUtils struct
type MethodUtils struct {
}

// Calcuator struct
type Calcuator struct {
	num1 int
	num2 int
}

func (c Calcuator) getRes(n1 int, n2 int, oper string) int {
	switch oper {
	case "+":
		return n1 + n2
	case "-":
		return n1 - n2
	case "*":
		return n1 * n2
	case "/":
		return n1 / n2
	default:
		fmt.Println("输入有误，请重新输入!")
		return -1
	}
}

func (m MethodUtils) myPrint() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func (m MethodUtils) print(n1, n2 int) {
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func (m MethodUtils) area(len, width float64) float64 {
	return len * width
}

func (m MethodUtils) judg(n int) {
	if n%2 == 0 {
		fmt.Printf("%v是偶数\n", n)
	} else {
		fmt.Printf("%v是奇数\n", n)
	}
}
func (m MethodUtils) print2(n1 int, n2 int, oper string) {
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			fmt.Printf(oper)
		}
		fmt.Println()
	}
}

func (m MethodUtils) print3(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%vX%v=%v\t", j, i, j*i)
		}
		fmt.Println()
	}
}

func (m MethodUtils) transpose(arr [3][3]int) [3][3]int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if i < j {
				arr[i][j], arr[j][i] = arr[j][i], arr[i][j]
			}
		}
	}
	return arr
}

func main() {
	var m MethodUtils
	m.myPrint()
	m.print(3, 4)
	res := m.area(4.0, 5.0)
	fmt.Println("面积 =", res)
	m.judg(345)
	m.print2(4, 4, "@")

	var c Calcuator
	result := c.getRes(3, 4, "*")
	fmt.Println("结果是", result)

	var n int
	fmt.Scanln(&n)
	m.print3(n)

	var arr = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	arr1 := m.transpose(arr)
	for _, v := range arr1 {
		for _, v2 := range v {
			fmt.Printf("%v ", v2)
		}
		fmt.Println()
	}

}
