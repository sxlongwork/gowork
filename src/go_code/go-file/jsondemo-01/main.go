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

// 将结构体实例序列化
func jsonStruct(){
	// 创建结构体实例
	user := User{"tom", 21, "beijing"}
	// json.Marshal()接受一个interface{}类型参数，即任意类型参数，返回[]byte, error
	data, err := json.Marshal(user)
	if err != nil{
		fmt.Println("json err, err =", err)
	}
	fmt.Println(string(data))
}

// 将map类型变量序列化
func jsonMap(){
	var m map[string]User
	m = make(map[string]User)
	m["001"] = User{"jack", 22, "shandong"}
	m["002"] = User{"mary", 23, "hebei"}

	data, err := json.Marshal(m)
	if err != nil{
		fmt.Println("json err, err =", err)
	}
	fmt.Println(string(data))
}

// 将切片类型实例序列化
func jsonSlice(){

	var slice []User
	user1 := User{"xiaoming", 23, "anhui"}
	user2 := User{"xianqian", 26, "anjing"}
	slice = append(slice, user1,user2)

	data, err := json.Marshal(slice)
	if err != nil{
		fmt.Println("json err, err =", err)
	}
	fmt.Println(string(data))
}

func main(){
	jsonStruct()
	jsonMap()
	jsonSlice()
}