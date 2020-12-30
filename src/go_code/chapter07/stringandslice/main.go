package main
import "fmt"

func main(){

	var str = "xianqian"
	s1 := str[2:5]
	fmt.Printf("s1 = %v, len(s1) = %v\n", s1, len(s1))

	//如果需要修改字符串，可以先将string -> []byte/或者 []rune -> 修改 -> 重写转成string
	// slice := []byte(str)
	// slice[0] = 'q'
	// str = string(slice)
	// fmt.Printf("str=%v", str)	// "qianqian"

	//注意以上方法处理字母或数字的字符串可以，但如果有汉字，会出现乱码，因为数字占3个字节，需要使用[]rune
	slice := []rune(str)
	slice[0] = '爱'
	str = string(slice)
	fmt.Printf("str=%v", str)	// "爱ianqian"
}