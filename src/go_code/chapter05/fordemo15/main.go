package main

import "fmt"

func main() {
	//先易后难
	//先死后活
	var class int = 3
	var stu int = 5
	var allAvg float64
	for i := 1; i <= class; i++ {
		avg := 0.0
		sum := 0.0
		for j := 1; j <= stu; j++ {
			var scole float64
			fmt.Println("请输入学生成绩：")
			fmt.Scanln(&scole)
			sum += scole
		}
		allAvg += sum
		avg = sum / float64(stu)
		fmt.Println("班级平均分是", avg)
	}
	fmt.Println("总平均分是", allAvg / float64(stu * class))

	
}