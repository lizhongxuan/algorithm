//题目:猫狗队列
package main

import (
	time2 "time"
	"sync"
)

type Cat struct {
	name string
	time int64
}

type Dog struct {
	name string
	time int64
}



type catDogQueue struct {
	catQueue *Queue
	dogQueue *Queue
	sync.Mutex
}


func NewCatDogQueue()*catDogQueue  {
	return &catDogQueue{
		catQueue:NewQueue(),
		dogQueue:NewQueue(),
	}
}

func NewCat(name string)*Cat  {
	return &Cat{
		name:name,
	}
}

func NewDog(name string)*Dog  {
	return &Dog{
		name:name,
	}
}


func (cdq *catDogQueue)Add(v interface{})bool  {
	cdq.Lock()
	defer cdq.Unlock()
	c,ok := v.(*Cat)
	if ok {
		time := time2.Now().UnixNano()
		c.time = time
		cdq.catQueue.Push(c)
		return true
	}

	d,ok := v.(*Dog)
	if ok {
		time := time2.Now().UnixNano()
		d.time = time
		cdq.dogQueue.Push(d)
		return true
	}

	return false
}

func (cdq *catDogQueue)PollAll()[]interface{}  {

	var a []interface{}
	for !cdq.IsEmpty() {
		
		if cdq.IsDogEmpty() {
			v := cdq.catQueue.Pop()
			a = append(a, v)
			continue
		}else if cdq.IsCatEmpty(){
			v := cdq.dogQueue.Pop()
			a = append(a, v)
			continue
		}

		if cdq.dogQueue.botton.value.(*Dog).time >= cdq.catQueue.botton.value.(*Cat).time {
			v := cdq.catQueue.Pop()
			a = append(a, v)

		}else {
			v := cdq.dogQueue.Pop()
			a = append(a, v)

		}
		
		
	}
	return a
}

func (cdq *catDogQueue)IsEmpty()bool{
	if cdq.catQueue.top == nil && cdq.dogQueue.top == nil {
		return true
	}
	return false
}
func (cdq *catDogQueue)IsDogEmpty()bool{
	if cdq.dogQueue.top == nil {
		return true
	}
	return false
}
func (cdq *catDogQueue)IsCatEmpty()bool{
	if cdq.catQueue.top == nil {
		return true
	}
	return false
}

func (cdq *catDogQueue)PollCat()[]interface{}{
	var a []interface{}
	for !cdq.IsCatEmpty() {
		v := cdq.catQueue.Pop()
		a = append(a, v)
	}
	return a
}
func (cdq *catDogQueue)PollDog()[]interface{}{
	var a []interface{}
	for !cdq.IsDogEmpty() {
		v := cdq.dogQueue.Pop()
		a = append(a, v)
	}
	return a
}

