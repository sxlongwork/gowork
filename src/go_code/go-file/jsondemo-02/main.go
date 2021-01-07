package main

import (
	"fmt"
	"encoding/json"
)

type User struct{
	Name string	`json:"name"`	// 反射机制
	Age int	`json:"age"`
	Address string	`json:"address"`
}

// 将结构体类型的json串反序列化为对应的结构体数据类型
func unJsonStruct(){
	var str = "{\"name\":\"tom\",\"age\":21,\"address\":\"beijing\"}"

	var user = User{}
	err := json.Unmarshal([]byte(str), &user)
	if err != nil{
		fmt.Println("Unmarshal failed, err =", err)
	}
	fmt.Println(user)
}
// 将map类型的json字符串反序列化为对应的map类型
func unJsonMap(){
	var str = "{\"001\":{\"Name\":\"jack\",\"Age\":22,\"Address\":\"shandong\"},\"002\":{\"Name\":\"mary\",\"Age\":23,\"Address\":\"hebei\"}}"
	var m map[string]User

	// 反序列化时，m不需要make,因为Unmarshal函数内部会make
	err := json.Unmarshal([]byte(str), &m)
	if err != nil{
		fmt.Println("Unmarshal failed, err =", err)
	}
	fmt.Println(m)
}

// 将切片类型的json字符串反序列化为对应的切片类型
func unJsonSlice(){
	var str = "[{\"Name\":\"xiaoming\",\"Age\":23,\"Address\":\"anhui\"},{\"Name\":\"xianqian\",\"Age\":26,\"Address\":\"anjing\"}]"
	var slice []User
	// 反序列化时，slice不需要make,因为Unmarshal函数内部会make
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil{
		fmt.Println("Unmarshal failed, err =", err)
	}
	fmt.Println(slice)
}


func main(){
	unJsonStruct()
	unJsonMap()
	unJsonSlice()
}