package user

import (
	"testing"
	"os"
)

func TestStore(t *testing.T){
	filePath := "./test.txt"
	var slice = []User{
		{
			Name: "tom",
			Age: 21,
			Address: "beijing",
		},
		{
			Name: "jack",
			Age: 22,
			Address: "xian",
		},
	}
	for _, v := range slice{
		v.Store()
		info, err := os.Stat(filePath)
		if err != nil{
			t.Fatalf("测试失败, 期望=读取文件成功, 实际=读取文件失败")
		}
		if info.Size() == int64(0){
			t.Fatalf("测试失败, 期望=文件大小不为0, 实际=文件大小为0")
		}
	}
}

func TestReStore(t *testing.T){
	user := User{}

	user.ReStore()
	if user.Name == "" || user.Age == 0 || user.Address == "" {
		t.Fatalf("测试失败, 期望=序列化后对象属性不为默认值, 实际=序列化后对象属性为默认值")
	}
}