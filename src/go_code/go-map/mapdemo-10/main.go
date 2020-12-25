package main

import "fmt"

func modifyUser(user map[string]map[string]string, name string) {
	if v, ok := user[name]; ok {
		v["pwd"] = "888888"
	} else {
		user[name] = make(map[string]string)
		user[name]["nickname"] = "hello"
		user[name]["pwd"] = "111111"
	}
}

func main() {
	var m = make(map[string]map[string]string)
	m["xiaoming"] = make(map[string]string)
	m["xiaoming"]["nickname"] = "xm"
	m["xiaoming"]["pwd"] = "123456"

	m["tom"] = make(map[string]string)
	m["tom"]["nickname"] = "tt"
	m["tom"]["pwd"] = "0923"

	modifyUser(m, "jack")
	for k, v := range m {
		fmt.Println(k)
		for k2, v2 := range v {
			fmt.Printf("\t%v=%v\n", k2, v2)
		}
	}
	// fmt.Println(m)
}
