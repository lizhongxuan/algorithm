//题目:由两个栈组成的队列

package main

import (
	"sync"
)

type stackToList struct {
	firstStack *Stack
	secondStack *Stack
	sync.Mutex
}

func NewStackToList() *stackToList {
	return &stackToList{
		firstStack:NewStack(),
		secondStack:NewStack(),
	}
}

func (sl *stackToList)Push(v interface{})  {
	sl.Lock()
	defer sl.Unlock()
	if sl.firstStack.top == nil &&  sl.secondStack.top == nil{
		sl.firstStack.Push(v)
		return
	}

	for sl.firstStack.top != nil  {
		value := sl.firstStack.Pop()
		sl.secondStack.Push(value)
	}
	sl.secondStack.Push(v)

	for sl.secondStack.top != nil  {
		value := sl.secondStack.Pop()
		sl.firstStack.Push(value)
	}

}

func (sl *stackToList)Pop() interface{} {
	sl.Lock()
	defer sl.Unlock()


	if sl.firstStack.top != nil  {
		value := sl.firstStack.Pop()
		return value
	}

	return nil
}