//题目:设计一个有getMin功能的栈

package main

import (
	"errors"
	"sync"
	"fmt"
)

type getMinStack struct {
	normalStack *Stack
	minStack *Stack
	sync.RWMutex
}

func NewGetMinStack()*getMinStack  {
	return &getMinStack{
		normalStack:NewStack(),
		minStack:NewStack(),
	}
}

func (gs *getMinStack)Push(v int64)bool  {
	gs.Lock()
	defer gs.Unlock()

	gs.normalStack.Push(v)

	if gs.minStack.top == nil || gs.minStack.top.value.(int64) >= v {
		fmt.Println("Push")
		gs.minStack.Push(v)
	}


	return true
}


func (gs *getMinStack)Pop() (int64,error)  {
	gs.Lock()
	defer gs.Unlock()

	if gs.normalStack.top == nil {
		return 0,errors.New("top is nil")
	}

	i := gs.normalStack.Pop()
	if i == gs.minStack.top.value{
		gs.minStack.Pop()
	}
	return i.(int64),nil
}

func (gs *getMinStack)GetMin() (int64,error)   {
	gs.RLock()
	defer gs.RUnlock()

	if gs.minStack.top == nil {
		return 0,errors.New("top is nil")
	}

	return gs.minStack.top.value.(int64),nil

}