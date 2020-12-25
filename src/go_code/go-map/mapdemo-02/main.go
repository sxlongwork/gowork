package main

import "fmt"

func main() {

	var m1 map[bool]string
	// m1[true] = "ok"	//panic: assignment to entry in nil map，需要make初始化后才能赋值使用
	m1 = make(map[bool]string)
	m1[true] = "ok"
	m1[false] = "oh,my god"

	fmt.Println(m1)
	s1 := make([]*string, 2)
	for k, v := range m1 {
		fmt.Printf("key=%v\tvalue=%v\tkey地址=%p\tvalue地址=%p\n", k, v, &k, &v) //map中每个key的地址都相同，每个value的地址也是相同的，这两块块地址是临时分配的
		s1 = append(s1, &v)
	}
	fmt.Println(s1)
}
