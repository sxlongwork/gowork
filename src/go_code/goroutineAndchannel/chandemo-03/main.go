package main

import "fmt"

type Cat struct{
	Name string	`json:"name"`
	Age int		`json:"age"`
}

func main(){
	intChan := make(chan int, 3)
	// 存放3个int类型数据
	intChan <- -5
	intChan <- 10
	intChan <- 200
	// intChan容量为3，再放就会报错
	// 取出数据

	n1 := <- intChan
	n2 := <- intChan
	n3 := <- intChan
	// intChan已经为空，再取就会报错
	fmt.Printf("n1 = %v, n2 = %v, n3 = %v\n", n1, n2, n3)

	mapChan := make(chan map[string]string, 10)
	m1 := make(map[string]string, 10)
	m1["name"] = "xiaoming"
	m1["mobile"] = "13200001111"

	m2 := make(map[string]string, 10)
	m2["city1"] = "beijing"
	m2["city2"] = "shanghai"

	mapChan <- m1
	mapChan <- m2

	m3 := <- mapChan
	m4 := <- mapChan
	fmt.Printf("m3 = %v, m4 = %v\n", m3, m4)

	catChan := make(chan Cat, 10)
	cat1 := Cat{
		Name: "tom",
		Age: 3,
	}
	cat2 := Cat{
		Name: "jony",
		Age: 4,
	}
	catChan <- cat1
	catChan <- cat2

	cat11 := <- catChan
	cat22 := <- catChan
	fmt.Printf("cat11 = %v, cat22 = %v\n", cat11, cat22)

	catChan2 := make(chan *Cat, 10)
	cat3 := Cat{Name: "tom~~", Age: 5}
	cat4 := Cat{Name: "jony~~", Age: 6}
	
	catChan2 <- &cat3
	catChan2 <- &cat4

	cat33 := <- catChan2
	cat44 := <- catChan2
	fmt.Printf("cat33 = %v, cat44 = %v\n", cat33, cat44)

	allChan := make(chan interface{}, 10)
	allChan <- 10
	allChan <- "xiaoming"
	allChan <- true
	allChan <- Cat{Name: "barry", Age: 5}

	v1 := <- allChan
	v2 := <- allChan
	v3 := <- allChan
	v4 := <- allChan
	fmt.Printf("v1=%v, v2=%v, v3=%v, v4=%v\n", v1, v2 ,v3, v4)
	// 注意，上面读出的v4下面的用法错误，编译不通过, 需要使用类型断言
	// fmt.Printf("v4.Name = %v\n", v4.Name)
	v5 := v4.(Cat)
	fmt.Printf("v5.Name = %v\n", v5.Name)
}