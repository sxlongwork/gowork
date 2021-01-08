package main
import (
	"fmt"
	"time"
	"sync"
)

var (
	// 定义全局map变量，用于存储数据
	m map[int]int = make(map[int]int, 10)
	// 定义锁变量,Mutex是互斥的意思，lock即全局的互斥锁
	lock sync.Mutex
)
func test(n int){
	res := 0
	for i := 1; i<=n; i++{
		res += i
	}
	// 当goroutine访问共享变量m时加上锁，不让其他routine访问
	lock.Lock()
	m[n] = res
	// 当goroutine访问共享变量m结束时放开锁
	lock.Unlock()
}

func main(){
	// 开启200个goroutine完成这个任务
	for i := 1; i <= 200; i++ {
		go test(i)
	}
	// 休眠10s等待goroutine执行完成【问题2: 不知道等多久?】
	time.Sleep(10 * time.Second)

	// 【问题1】fatal error: concurrent map writes 资源竞争错误
	lock.Lock()
	for i, v := range m{
		fmt.Printf("m[%v] = %v\n", i, v)
	}
	lock.Unlock()
}