package main

import (
	"fmt"
	"sort"
)

/*
贪心算法
在对问题求解时，总是做出在当前看来是最好的选择。也就是说，不从整体最优上加以考虑，算法得到的是在某种意义上的局部最优解
贪心算法一般按如下步骤进行
①建立数学模型来描述问题。
②把求解的问题分成若干个子问题。
③对每个子问题求解，得到子问题的局部最优解 。
④把子问题的解局部最优解合成原来解问题的一个解
*/

/*
题目描述：每个孩子都有一个满足度，每个饼干都有一个大小，只有饼干的大小大于一个孩子的满足度，该孩子才会获得满足。求解最多可以获得满足的孩子数量
c[1,2,3] 孩子的满足度; s[1,1,3] 饼干的大小
求解可以满足孩子的数量
*/
func getChildren(c []int, s []int) int {
	sort.Ints(c)
	sort.Ints(s)
	i := 0
	j := 0
	count := 0
	for i < len(c) && j < len(s) {
		if c[i] <= s[j] {
			count++
			i++
		}
		j++
	}
	return count
}

func main() {
	var c = []int{1, 2, 3, 4}
	var s = []int{1, 1, 2, 3}
	res := getChildren(c, s)
	fmt.Println(res)
}
