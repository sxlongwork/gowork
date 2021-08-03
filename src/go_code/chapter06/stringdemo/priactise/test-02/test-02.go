package main

import (
	"math"
	"fmt"
	"strconv"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func NewListNode(n int) *ListNode{
	root := new(ListNode)
	root.Val = n
	return root
} 
func (ln *ListNode) append(n int)  {
	cur := ln
	for {
		if cur.Next == nil{
			cur.Next = NewListNode(n)
			// fmt.Println(cur.Val)
			break
		} else {
			cur = cur.Next
		}
	}
}

func listNodeToSlice(listNode *ListNode) []int{
	var list1 []int = make([]int, 0, 10)
	for  {
		if listNode != nil{
			// fmt.Println(l1)
			list1 = append(list1, listNode.Val)
			listNode = listNode.Next
		} else {
			break
		}
	}
	return list1
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) []int {

	
	var list1 = listNodeToSlice(l1)
	var list2 = listNodeToSlice(l2)

	var num1 int = 0
	var num2 int =0 
	// fmt.Println(list1)
	
	for i, v := range list1{
		num1 += (v * int(math.Pow10(i)))
	}
	for i, v := range list2{
		num2 += (v * int(math.Pow10(i)))
	}
	result := num1 + num2
	var resultList = []rune(strconv.Itoa(result))

	var root *ListNode
	for i:=0 ;i < len(resultList) ; i++{
		f := fmt.Sprintf("%c", resultList[len(resultList) - i - 1])
		d, _ := strconv.Atoi(f)
		// fmt.Printf("%T,%d", d,d)
		if root == nil{
			root = NewListNode(d)
			// fmt.Println(root.Val)
		} else {
			root.append(d)
		}
	}
	var list3 = listNodeToSlice(root)
	return list3
}


func main(){
	var l1 *ListNode = NewListNode(2)
	l1.append(4)
	l1.append(3)

	var l2 *ListNode = NewListNode(5)
	l2.append(6)
	l2.append(4)

	fmt.Println(addTwoNumbers(l1,l2))

}