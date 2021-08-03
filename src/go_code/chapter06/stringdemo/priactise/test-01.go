package main

import (
	"bufio"
	"strings"
	"strconv"
	"sort"
	"os"
	"fmt"
)

/*
去除数组重复元素
输入：1 1 1 1 3 4 5 3
输出：1 3 4 5 
*/
func test01() {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	datalist := strings.Fields(string(data))

	var numlist []int = make([]int, 0,10)
	for _, v := range datalist {
		num, _ := strconv.Atoi(v)
		numlist = append(numlist, num)
	}
	sort.Sort(sort.IntSlice(numlist))

	for i := 0; i< len(numlist)-1; i++ {
		if numlist[i] ^ numlist[i+1] == 0 {
			numlist = append(numlist[:i], numlist[i+1:]...)
			i--
		}
	}
	for _, v := range numlist{
		fmt.Print(v, " ")
	}
}
/*
计算字符个数
输入:		aAbcf
			a
输出:		2
*/
func test02(){
	reader := bufio.NewReader(os.Stdin)

	str, _ := reader.ReadString('\n')
	ch, _ := reader.ReadByte()
	num := strings.Count(strings.ToLower(str), strings.ToLower(string(ch)))
	fmt.Println(num)
}

/*
明明的随机数
*/
func myMing(){
	reader := bufio.NewReader(os.Stdin)
	// count := 0
	sample := make([]int,0,10)
	numlist := make([]int,0,10)

	for num ,_ ,_ := reader.ReadLine(); len(num) != 0; {
		n, _ := strconv.Atoi(string(num))
		
		for i :=0; i< n; i++ {
			data, _, _ :=  reader.ReadLine()
			intn,_ := strconv.Atoi(string(data))
			numlist = append(numlist, intn)
		}
		// sample[count] = n
		sample = append(sample, n)
		// count++
		num ,_ ,_ = reader.ReadLine()
	}
	outlist := make([]int,0,10)
	for _,v := range sample{
		outlist = append(outlist, numlist[:v]...)
		sort.Sort(sort.IntSlice(outlist))

		for i := 0; i< len(outlist)-1; i++ {
			if outlist[i] ^ outlist[i+1] == 0 {
				outlist = append(outlist[:i], outlist[i+1:]...)
				i--
			}
		}
		for _, val := range outlist{
			fmt.Println(val)
		}
		
		numlist = numlist[v:]
		outlist = make([]int,0,10)

	}
}

/*
明明的随机数
*/
func test03(){
	scan := bufio.NewScanner(os.Stdin)
    for {
        scan.Scan()
        if len(scan.Text()) == 0 {
            break
        }
        count,_ := strconv.Atoi(scan.Text())
        n := [1001]bool{}
        for i := 0;i < count;i++ {
            scan.Scan()
            numstr := scan.Text()
            num,_ := strconv.Atoi(numstr)
            n[num] = true
        }
         
        for i,v := range n {
        if v == true {
            fmt.Println(i)
        }
    }
    }
}

func test04(){
	reader := bufio.NewReader(os.Stdin)
	var nums [1001]bool

	for {
		data , _, _ := reader.ReadLine()
		if len(data) == 0{
			break
		}
		n, _ := strconv.Atoi(string(data))
		for i:=0; i< n; i++{
			numstr, _, _ := reader.ReadLine()
			num, _ := strconv.Atoi(string(numstr))
			nums[num] = true
		}
	}
	for i, v := range nums{
		if v == true{
			fmt.Println(i)
		}
	}
}

func main(){
	test03()
	
}