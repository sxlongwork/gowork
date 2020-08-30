package main

// import "fmt"
// import "unsafe"
import (
	"fmt"
	"unsafe"
)

func main() {
	
	//byte等价于unint8,可以用来表示单个字母字符，0-255
	var i byte = 255
	fmt.Println("i =", i)

	var n int32 = 10
	fmt.Printf("n数据类型是%T, n占用字节大小是%d", n, unsafe.Sizeof(n))
}