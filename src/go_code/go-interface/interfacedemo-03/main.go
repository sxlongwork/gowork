package main
import (
	"fmt"
	"math/rand"
	"sort"
)

type Person struct {
	Name string
	Age int
}

type Pslice []Person

//Len方法返回集合中的元素个数
func (ps Pslice) Len() int{
	return len(ps)
}

// Less方法报告索引i的元素是否比索引j的元素小
func (ps Pslice) Less(i, j int) bool{
	return ps[i].Age < ps[j].Age
}

// Swap方法交换索引i和j的两个元素
func (ps Pslice) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func main(){
	var ps Pslice = make([]Person,10)	//注意这里声明时必须写上Pslice，不然就是[]Person类型，不是Pslice类型了
	//随机生成10个学生信息
	for i := 0; i < 10; i++{
		name := fmt.Sprintf("stu-%d",rand.Intn(100))
		age := rand.Intn(100)
		ps[i] = Person{name, age}
	}
	fmt.Println("排序前")
	for _, v := range ps{
		fmt.Println(v)
	}
	sort.Sort(ps)
	fmt.Println("排序后")
	for _, v := range ps{
		fmt.Println(v)
	}


}