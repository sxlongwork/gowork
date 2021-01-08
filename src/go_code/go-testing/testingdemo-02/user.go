package user

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type User struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Address string `json:"address"`
}

func (user *User) Store(){
	// 将对象序列化
	data, err := json.Marshal(*user)
	if err != nil{
		fmt.Println("json user error, err =", err)
	}

	// 将data写入文件中，可以使用ioutil.WriteFile(file,data)
	filePath := "./test.txt"
	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil{
		fmt.Println("write to file error, err =", err)
	}
}

func (user *User) ReStore(){
	// 读取文件
	filePath := "./test.txt"
	data, err := ioutil.ReadFile(filePath)
	if err != nil{
		fmt.Println("read file error, err =", err)
	}

	// 将数据反序列化并存入对应的变量中
	err = json.Unmarshal(data, user)
	if err != nil{
		fmt.Println("Unmarshal error, err =", err)
	}
}