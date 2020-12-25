package main

import "fmt"

func test(m map[string]string, key string) {
	delete(m, key)
}

func main() {
	var m = map[string]string{
		"name": "tom",
		"age":  "20",
	}
	fmt.Println("调用test前:", m)
	test(m, "age")
	fmt.Println("调用test后:", m) //调用test后，原来的m改变了，说明map是引用传递

	var m1 = make(map[string]string, 2) //定义容量为2
	m1["01"] = "aa"
	m1["02"] = "bb"
	m1["03"] = "cc" //不会报错，会自动扩容
	fmt.Println(m1)
}
