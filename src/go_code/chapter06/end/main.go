package main
import (
	"fmt"
)

func isLeapYear(year int) bool {
	if (year % 400 == 0) || (year % 4 == 0 && year % 100 != 0) {
		return true
	} else {
		return false
	}	
}

func getDays(year int, month int) {

	isleapyear := isLeapYear(year)
	switch(month) {
	case 1,3,5,7,8,10,12:
		fmt.Printf("%v年%v月共有31天\n",year,month)
	case 4,6,9,11:
		fmt.Printf("%v年%v月共有30天\n",year,month)
	case 2:
		if isleapyear {
			fmt.Printf("%v年%v月共有29天\n",year,month)
		} else {
			fmt.Printf("%v年%v月共有28天\n",year,month)
		}
	default:
		fmt.Println("输入的月份有误，请重新输入")
		// continue
	}
}

func main(){
	// fmt.Println(isLeapYear(2100))

	var year, month int
	for {
		fmt.Printf("请输入年份：")
		fmt.Scanln(&year)
		fmt.Printf("请输入月份：")
		fmt.Scanln(&month)

		getDays(year, month)
	}
	
}