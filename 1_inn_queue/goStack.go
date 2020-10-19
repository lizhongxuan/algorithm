//golang stack realize

package main

import "fmt"

type Stack struct {
	top *stackNode
	botton *stackNode
}

type stackNode struct {
	value interface{}
	nextNode *stackNode
	preNode *stackNode
}

func NewStack()*Stack  {
	return &Stack{
		top:nil,
		botton:nil,
	}
}


func (s *Stack)Push(value interface{})bool  {
	if s.top == nil {
		node := &stackNode{
			value:value,
		}
		s.top = node
		s.botton = node
		return true
	}

	node := &stackNode{
		value:value,
	}

	s.top.preNode = node
	node.nextNode = s.top
	s.top = node

	return true
}

func (s *Stack)Pop()interface{}  {

	if s.top == nil  {
		return nil
	}
	node := s.top

	if s.top.nextNode == nil {
		s.top = nil
		s.botton = nil
	}else {
		s.top = s.top.nextNode
		s.top.preNode = nil
	}

	return node.value

}

func (s *Stack)readAllStack()  {
	fmt.Println("begin readAllStack !!!")

	if s.top == nil {
		fmt.Println("stack is nil")
		fmt.Println("end ----------->")
		return
	}

	node := s.top
	fmt.Println("node value :",node.value)

	for node.nextNode != nil  {
		node = node.nextNode
		fmt.Println("node value :",node.value)
	}

	fmt.Println("end ----------->")
}


func (s *Stack)isEmpty()bool  {
	if s.top == nil {
		return true
	}
	return false
}

func (s *Stack)Len()int  {
	i := 0
	node := s.top
	for node != nil {
		node = node.nextNode
		i++
	}
	return i
}