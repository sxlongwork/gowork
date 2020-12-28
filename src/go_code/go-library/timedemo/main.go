package main

import (
	"fmt"
	"time"
)

func main() {
	//time包的使用

	//1. 获取当前时间, t为time.Time类型，Time是一个struct，返回了一个Time的实例，根据该实例就可以调用其属性或方法
	t := time.Now() //2020-12-28 14:26:15.8984944 +0800 CST m=+0.006977501
	fmt.Println(t)

	//2. 获取年月日时分秒
	year := t.Year()
	month := int(t.Month())
	day := t.Day()
	hour := t.Hour()
	minute := t.Minute()
	sec := t.Second()
	fmt.Println(year, month, day, hour, minute, sec)

	//3. 时间字符串转化为Time类型,Parse方法转换的结果是UTC时区的时间，会与time.Now（CST）有8小时的差值;
	// t1, error := time.Parse("2006/1/2 15:04:05", "2020/12/28 14:44:12") //UTC时间
	t1, error := time.ParseInLocation("2006/1/2 15:04:05", "2020/12/28 14:44:12", time.Local) //time.ParseInLocation带时区的转化
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(t1)

	//4. 时间的格式化
	// str := t1.Format("2006-1-2 15:04:05")
	str := t1.Format("2006/1/2 15:04:05")
	fmt.Println(str)

	//5. 常用的时间常量
	// const (
	// 	Nanosecond  Duration = 1
	// 	Microsecond          = 1000 * Nanosecond
	// 	Millisecond          = 1000 * Microsecond
	// 	Second               = 1000 * Millisecond
	// 	Minute               = 60 * Second
	// 	Hour                 = 60 * Minute
	// )
	//休眠2小时=time.Sleep(2 * Hour);休眠20分钟=time.Sleep(20 * Minute);休眠30毫秒=time.Sleep(30 * Millisecond)

	//6. 获取当前unix时间戳和unixnano时间戳(作用时获取一些随机数字)
	fmt.Println(t1.Unix())
	fmt.Println(t1.UnixNano())
}
