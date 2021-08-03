/*
单链表是一种顺序存储的结构。 
有一个头结点，没有值域，只有连域，专门存放第一个结点的地址。 
有一个尾结点，有值域，也有链域，链域值始终为NULL。 
所以，在单链表中为找第i个结点或数据元素，必须先找到第i - 1 结点或数据元素，而且必须知道头结点，否者整个链表无法访问。

本文主要通过Golang实现链表的几种常见操作:
1、判断是否为空的单链表
2、单链表的长度
3、从头部添加元素
4、从尾部添加元素
5、在指定位置添加元素
6、删除指定元素
7、删除指定位置的元素
8、查看是否包含某个元素
9、遍历所有元素
*/

package linkedlist

type object interface{}

type ListNode struct {
	Val object
	Next *ListNode
}

type List struct {
	headNode *ListNode
}

func (list *List) IsEmpty() bool {
	if list.headNode == nil {
		return true
	}
	return false
}

func (list *List) GetLength() int {
	count := 0
	cur := list.headNode
	for cur != nil{
		count++
		cur = cur.Next
	}
	return count
}

func (list *List) Add(val object) {
	node := new(ListNode)
	node.Val = val
	node.Next = list.headNode
	list.headNode = node
}

func (list *List) Append(val object) {
	node := new(ListNode)
	node.Val = val
	if list.headNode == nil{
		list.headNode = node
	} else{
		cur := list.headNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
	
}

