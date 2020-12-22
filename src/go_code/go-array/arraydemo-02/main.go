package main

import "fmt"

func main() {
	var class = [3][5]float64{{1, 1, 1, 1, 1}, {2, 2, 2, 2, 2}, {3, 3, 3, 3, 3}}

	allScore := 0.0
	for i, v := range class {
		classScore := 0.0
		for _, v2 := range v {
			classScore += v2
		}
		fmt.Printf("第%v个班级的总分为%v, 班级平均分为%v\n", i, classScore, classScore/float64(len(class[i])))
		allScore += classScore
	}
	fmt.Printf("所有班级的总分为%v, 平均分为%v\n", allScore, allScore/float64(15))
}
