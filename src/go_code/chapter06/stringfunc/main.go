package main
import (
	"fmt"
	"strconv"
	"strings"
)

//常用的字符串系统函数

func main(){
	
	//统计字符串的长度，按字节len(str)
	str := "xianqian海"					//一个字母占一个字节，一个汉字占3个字节，所以共11个字节
	fmt.Println("len(str)=", len(str))	// 11

	//字符串遍历，同时处理中文问题，r = []rune(str)
	// r := []rune("xianqian么么哒")
	// for i :=0; i < len(r); i++ {
	// 	fmt.Printf("%c\n", r[i])
	// }

	//字符串转换为整数
	// n, error := strconv.Atoi("123么么哒")
	// if error != nil {
	// 	fmt.Println("转换错误", error)
	// } else {
	// 	fmt.Println("转换成功", n)
	// }

	//整数转字符串
	str = strconv.Itoa(123)
	fmt.Println(str)

	//字符串转[]byte，[]byte("hello")
	// b := []byte("xianqian")
	// fmt.Printf("b=%v\n",b)

	//byte[]转字符串，str = string([]byte{97,98,99,100})
	// str = string([]byte{97,98,99,100})
	// fmt.Println(str)

	//10进制转2，8，16进制，str = strconv.FormatInt(123, 2)
	// str = strconv.FormatInt(123, 2)
	// fmt.Println(str)
	// str = strconv.FormatInt(123, 8)
	// fmt.Println(str)
	// str = strconv.FormatInt(123,16)
	// fmt.Println(str)
	// str = strconv.FormatInt(123,10)
	// fmt.Println(str)

	//查找字符是否在指定的字符串中，f = strings.Contains(str, substr)
	// f := strings.Contains("helloxianqianlove", "xianqian")	//true
	// f := strings.Contains("helloxianqianlove", "lovexian")		//false
	// fmt.Println(f)

	//统计一个字符串中有几个指定的子串，strings.Count(str,substr)
	// n := strings.Count("helloxianqian", "n")
	// fmt.Println("n=", n)

	//不区分大小写的字符串比较(==是区分字母大小写的)，strings.EqualFold(str1, str2)
	// a := strings.EqualFold("abc", "AbC")
	// fmt.Println("结果是", a)				//true
	// fmt.Println("结果是", "abc"=="AbC")		//false

	//返回子串在字符串中第一次出现的index值，如果没有，就返回-1，strings.Index(str, substr)
	// index := strings.Index("lovexianqian", "xianqian")	//4
	// index := strings.Index("lovexianqian", "hello")			//-1
	// fmt.Println("index=", index)

	//返回子串在字符串中最后一次出现的index值，如果没有，就返回-1，strings.LastIndex(str, substr)
	// index := strings.LastIndex("lovexianqianlove", "love")	//12
	// index := strings.LastIndex("lovexianqianlove", "hello")		//-1
	// fmt.Println("index = ", index)

	//将指定的子串替换为另一个子串，strings.Replace(str, str1, str2, n),n表示替换几个，n=-1表示全部替换
	// str = strings.Replace("love xianqian love", "love", "love you", 1)	//love you xianqian love
	// str = strings.Replace("love xianqian love", "love", "love you", -1)		//love you xianqian love you
	// fmt.Println("str = ", str)

	//按照某个指定的字符，为分隔符，将字符串分割为一个字符串数组，strings.Split(str, substr)
	// strArr := strings.Split("love,xian,qian", ",")
	// for i:=0; i< len(strArr); i++ {
	// 	fmt.Printf("strArr[%v]=%v\n", i, strArr[i])
	// }

	//将字符串进行大小写转换，strings.ToLower("GO")，strings.ToUpper("go")
	// str = strings.ToLower("XIANQIAN")	//xianqian
	// str = strings.ToUpper("love")			//LOVE
	// fmt.Println(str)

	//将字符串左右两边的空格去掉，strings.TrimSpace(str)
	// str = strings.TrimSpace("   love   ")	//love
	// fmt.Println(str)

	//将字符串左右两边指定的字符去掉，strings.Trim("!hello!", "!")
	// str = strings.Trim("! hello !", " !")	// hello,前后的空格和!都被去掉，中间的去不掉
	// fmt.Println(str, len(str))

	//将字符串左边指定的字符去掉，strings.TrimLeft("!hello!")
	// str = strings.TrimLeft("!love!", "!")	//love!
	// fmt.Println(str)
	
	//将字符串右边指定的字符去掉，strings.TrimRight("!hello!")
	// str = strings.TrimRight("!love!", "!")		//!love
	// fmt.Println(str)

	//20、判断字符串是否以指定的字符串开头，strings.HasPrefix("lovexianqian", "love")
	f := strings.HasPrefix("lovexianqian", "love")	//true
	fmt.Println(f)

	//21、判断字符串是否以指定的字符串结束，strings.HasSuffix("xianqianlove", "love")
	f := strings.HasSuffix("xianqianlove", "love")		//true
	fmt.Println(f)
}