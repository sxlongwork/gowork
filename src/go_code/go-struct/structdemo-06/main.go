package main

import (
	"encoding/json"
	"fmt"
)

// Person struct
type Person struct {
	Age   byte    `json:"age"`
	Score float64 `json:"score"`
}

func main() {
	var p = Person{
		Age:   10,
		Score: 87.5,
	}
	fmt.Printf("age地址:%p\n", &p.Age) //Age是byte类型，占用1个字节，8位
	fmt.Printf("Score地址:%p\n", &p.Score)

	//将p变量序列化为json格式字串
	jsonStr, err := json.Marshal(p) //jsonStr是一个[]byte类型，需要转化为string字符串输出,string([]byte)
	if err != nil {
		fmt.Println("序列化失败!")
	}
	fmt.Println(string(jsonStr))
}
