package cal

import (
	"testing"
)

func TestAddUper(t *testing.T){
	var s1 = []int{0,5,10}
	var s2 =[]int{0,15,55}
	for index, _ := range s1{
		if res := AddUpper(s1[index]);  res != s2[index] {
			t.Fatalf("测试失败, 期望值=%v, 实际值=%v", s2[index], res)
		}
	}
}