package arraylist

import (
	"errors"
	"fmt"
)

type List interface {
	Size() int
	Get(index int) (interface{}, error)
	Update(index int, newVal interface{}) error
	Append(newVal interface{})
	Insert(index int, newVal interface{}) error
	Clear()
	Delete(index int) error
	String() string
}

type ArrayList struct {
	data    []interface{}
	theSize int
}

func NewArrayList() *ArrayList {
	list := new(ArrayList)
	list.data = make([]interface{}, 0, 10)
	list.theSize = 0
	return list
}

func (list *ArrayList) Size() int {
	return list.theSize
}

func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.theSize {
		return nil, errors.New("下标越界")
	}
	return list.data[index], nil
}

func (list *ArrayList) Update(index int, newVal interface{}) error {
	if index < 0 || index >= list.theSize {
		return errors.New("下标越界")
	}
	list.data[index] = newVal
	return nil
}

func (list *ArrayList) Append(newVal interface{}) {
	list.data = append(list.data, newVal)
	list.theSize++
}

func (list *ArrayList) checkIsFull() {
	if list.theSize == cap(list.data) {
		newData := make([]interface{}, 2*list.theSize, 2*list.theSize)
		copy(newData, list.data)
		list.data = newData
	}
}

func (list *ArrayList) Insert(index int, newVal interface{}) error {
	if index < 0 || index >= list.theSize {
		return errors.New("下标越界")
	}
	list.checkIsFull()
	list.data = list.data[:list.theSize+1]
	for i := list.theSize; i > index; i-- {
		list.data[i] = list.data[i-1]
	}
	list.data[index] = newVal
	list.theSize++
	return nil

}

func (list *ArrayList) Clear() {
	list.data = make([]interface{}, 0, 10)
	list.theSize = 0
}
func (list *ArrayList) Delete(index int) error {
	if index < 0 || index >= list.theSize {
		return errors.New("下标越界")
	}
	list.data = append(list.data[:index], list.data[index+1:]...)
	list.theSize--
	return nil
}
func (list *ArrayList) String() string {
	return fmt.Sprint(list.data)
}
