package main
import (
	"fmt"
	"time"
	"strconv"
)

func test03(){
	str := ""
	for i := 0; i < 100000; i++{
		str += "hello" + strconv.Itoa(i)
	}
}

func main(){
	//日期和时间相关函数
	//1 获取当前时间
	now := time.Now()
	fmt.Printf("now = %v\nnow数据类型 = %T\n", now, now)

	//获取当前时间的年月日，时分秒
	// fmt.Printf("年=%v, type=%T\n", now.Year(),now.Year())
	// fmt.Printf("月=%v, type=%T\n", now.Month(),now.Month())
	// fmt.Printf("月=%v\n", int(now.Month()))
	// fmt.Printf("日=%v, type=%T\n", now.Day(),now.Day())
	// fmt.Printf("时=%v, type=%T\n", now.Hour(),now.Hour())
	// fmt.Printf("分=%v, type=%T\n", now.Minute(),now.Minute())
	// fmt.Printf("秒=%v, type=%T\n", now.Second(), now.Second())

	//格式化时间日期,方式一
	// fmt.Printf("日期时间：%2d-%2d-%2d %2d:%2d:%d\n", now.Year(),int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	// dateStr := fmt.Sprintf("日期时间：%2d-%2d-%2d %2d:%2d:%d", now.Year(),int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	// fmt.Printf("dateStr =%v \n", dateStr)

	//格式化时间日期，方式2
	// fmt.Printf("date=%v\n", now.Format("2006-01-02 15:04:05"))	//数字时固定写法，不能修改，格式可以修改。如"2006/01/02 15:04:05"
	// fmt.Printf("date=%v\n", now.Format("2006/01/02"))			//可以只输出年月日
	// fmt.Printf("date=%v\n", now.Format("15:04:05"))				//可以指输出时分秒
	// fmt.Printf("date=%v\n", now.Format("2006"))					//可以只输出年
	// fmt.Printf("date=%v\n", now.Format("01"))					//可以只输出月

	//每隔1s打印一个数字
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println("i = ", i)
	// 	// time.Sleep(1 * time.Second)
	// 	time.Sleep(1000 * time.Millisecond)
	// }

	//获取当前unix时间戳和unixnano时间戳(作用时获取一些随机数字)
	// fmt.Println("unix时间戳=",now.Unix())
	// fmt.Println("unix时间戳纳秒=",now.UnixNano())

	//统计test03函数运行的时间
	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
	time := end - start
	fmt.Printf("test03函数运行时间为%d秒", time)
}