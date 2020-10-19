//题目:用栈来求解汉诺塔问题

package main

import (
	"fmt"
)

//一共只有四个动作,而只有一个动作不违反小压达和相邻不可逆原则,另外三个动作一定都会违反

const (
	hNo =1
	LToM =2
	MToL = 3
	MToR = 4
	RToM = 5
)

type Action int

func fStackTotStack(record []Action,preNoAct Action,nowAct Action,
						fStack *Stack,tStack *Stack,from string,to string) int {
	if fStack.top == nil  {
		return 0
	}

	if record[0] != preNoAct && (tStack.top == nil || fStack.top.value.(int) < tStack.top.value.(int) ){
		tStack.Push(fStack.Pop())
		fmt.Println("Move ",tStack.top.value.(int),"  from ",from,"  to ",to)
		record[0] = nowAct
		return 1
	}
	return 0
}


func hanoiProblem(num int,left string,mid string,right string) int{
	 lS := NewStack()
	 mS := NewStack()
	 rS := NewStack()
	for i:=num; i>0;i--  {
		lS.Push(i)
	}
	record := []Action{hNo}
	step := 0
	for rS.Len() != num {
		step += fStackTotStack(record,MToL,LToM,lS,mS,left,mid)
		step += fStackTotStack(record,LToM,MToL,mS,lS,mid,left)
		step += fStackTotStack(record,RToM,MToR,mS,rS,mid,right)
		step += fStackTotStack(record,MToR,RToM,rS,mS,right,mid)
		fmt.Println("   ")
	}

	return step
}




























