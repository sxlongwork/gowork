package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main1(){

	reader := bufio.NewReader(os.Stdin)
	// data, _, _ := reader.ReadLine()
	// dataStr := string(data)
	// fmt.Println(dataStr)

	str, _ := reader.ReadString('\n')
	strlist := strings.Fields(str)
	fmt.Println(len(strlist[len(strlist)-1]))

}

func main(){

	// reader := bufio.NewReader(os.Stdin)
	// str, _ := reader.ReadString('\n')
	// ch, _ := reader.ReadByte()
	// num := strings.Count(strings.ToLower(str), strings.ToLower(string(ch)))
	// fmt.Println(num)

	str1 := "a Ba号2"
	str2 := "A ba号2"
	fmt.Println(strings.EqualFold(str1, str2))	// 忽略大小写
	fmt.Println(strings.EqualFold("abc", "AbC"))	//true
	fmt.Println(strings.HasPrefix("#_abc123", "_#"))	//false
	fmt.Println(strings.HasSuffix("#_abc123", "3"))	//true
	fmt.Println(strings.Contains("abcfkg", "as"))	//false
	fmt.Println(strings.ContainsAny("#中_abc123", "中"))	//true
	fmt.Println(strings.ContainsRune("#中_abc123", rune(12)))	//false
	fmt.Println(strings.Count("aabbcccA1233","a"))	// 2, 区别大小写，要想忽略大小写统计，可以先全部转为大写或小写
	fmt.Println(strings.Index("aabb123", "cb"))	// -1, 返回-1表示不存在
	fmt.Println(strings.LastIndex("aaccbbccd", "cc"))	//6
	fmt.Println(strings.Title("apple, blue"))	// Apple, Blue
	fmt.Println(strings.ToTitle("apple, blue"))	// APPLE, BLUE
	fmt.Println(strings.ToLower("OH, MY GOD"))	// oh, my god
	fmt.Println(strings.ToUpper("oh, my god"))	//OH, MY GOD
	fmt.Println(strings.Repeat("xy", 3))	// 重复3次xy
	fmt.Println(strings.Replace("hel hel hel ", "l", "llo", -1)) // 字符串替换
	fmt.Println(strings.Trim("aaaabbbaaaaa", "a"))	// bbb, 去掉两端包含a的所有字符串
	fmt.Println(strings.TrimSpace("   bbcc zz     "))	// bbcc zz, 去掉两端所有空白
	fmt.Println(strings.TrimLeft("aab    abcaab", "aab"))
	fmt.Println(strings.TrimRight("aab    abcaab", "aab"))
	fmt.Println(strings.TrimPrefix("aab    abcaab", "aab"))
	fmt.Println(strings.TrimSuffix("aab    abcaab", "aab"))
	fmt.Println(strings.Fields("aabb   ccdf"))	// 按空白(可以是多个连续空白字符)分隔返回字符串切片
	fmt.Println(strings.Split("aa,bb,cc", ","))
	str := make([]string, 0, 10)
	str = append(str, "aa")
	str = append(str, "bb")
	fmt.Println(strings.Join(str, "#"))	// aa#bb

}