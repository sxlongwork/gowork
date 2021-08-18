package main

import (
	"fmt"
)

/*
	中心扩展方式
*/
func test() {
	test := "abcde"
	newStr := make([]string, 0, 10)
	for i := 0; i < len(test); i++ {
		newStr = append(newStr, "#")
		newStr = append(newStr, test[i:i+1])
	}
	newStr = append(newStr, "#")
	fmt.Println(newStr)
	// for _, v := range newStr {
	// 	fmt.Printf("%v, %T\n", v, v)
	// }
	max_len := 0
	for i := 0; i < len(newStr); i++ {
		//获取以第i个字符为中心的回文串的长度
		cur_len := getMaxLen(newStr, i)
		if cur_len > max_len {
			max_len = cur_len
		}
	}
	fmt.Println(max_len)

}

/*
中心扩展函数
*/
func getMaxLen(data []string, i int) int {
	cur_len := 0
	for k := 0; k <= i && k <= len(data)-i-1; k++ {
		if data[i-k] == data[i+k] {
			cur_len++
		} else {
			break
		}
	}
	return cur_len - 1
}

func main() {
	test()
}
