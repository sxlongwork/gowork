// typedef struct sta
// {
//     int *top;            /* 栈顶指针 */
//     int *bottom;        /* 栈底指针 */
//     int stack_size;        /* 栈的最大容量 */
// }stack;
// stack Push (stack p);        /* 入栈 */
// void DisplyStack (stack p);    /* 遍历栈中元素 */
// stack Pop (stack p);        /* 出栈 */
// stack InitStack (stack p);    /* 初始化栈 */
// int StackEmpty (stack p);    /* 判断栈是否为空 */
// int StackFull (stack p);    /* 判断栈是否为满 */
package stack

import (
	"errors"
	"fmt"
)
type Stack interface{
	Size() int
	Push(newVal interface{})
	Pop() interface{}
	IsEmptyStack() bool
	IsFullStack()
	String()
	
	// Clear()
}

type MyStack struct{
	bottom int
	top int
	data []interface{}
	size int
}

func (stack *MyStack) Size() int {
	return stack.size
}

func NewStack() *MyStack{
	mystack := new(MyStack)
	mystack.bottom = 0
	mystack.top = 0
	mystack.size = 0
	mystack.data = make([]interface{},0,10)
	return mystack
}

func (stack *MyStack) Push(newVal interface{}) {
	stack.IsFullStack()
	// stack.data = stack.data[:stack.size+1]
	stack.data = append(stack.data, newVal)
	stack.top++
	stack.size++
}

func (stack *MyStack) IsFullStack(){
	if stack.size == cap(stack.data){
		newData := make([]interface{}, stack.size, 2*stack.size)
		copy(newData, stack.data)
		stack.data = newData
	}
}

func (stack *MyStack) Pop() interface{}{
	
	if stack.IsEmptyStack() != nil {
		return nil
	}

	val := stack.data[stack.top-1]
	stack.data = append(stack.data[:stack.top-1], stack.data[stack.top:]...)
	stack.top--
	stack.size--
	return val
}
func (stack *MyStack) IsEmptyStack() error  {
	if stack.top == stack.bottom {
		return errors.New("stack is empty")
	}
	return nil
}

func (stack *MyStack) String() string{
	return fmt.Sprint(stack.data)
}
