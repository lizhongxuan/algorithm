package main

import "fmt"

func main() {
	test9()

}

func test9()  {
	//题目:构造数组的MaxTree
	//maxTree.go
}

func test8()  {
	//生成窗口最大值数组
	//windowMaxArr.go
	arr := []int{1,1,1,3,3,3,4,4,4,1}
	fmt.Println(getMaxWindow(arr,3))
}

func test7()  {
	//题目:用栈来求解汉诺塔问题
	//hanoi.go
	i := hanoiProblem(3,"leftStack","midStack","rightStack")
	fmt.Println("step :",i)
}

func test6()  {
	//题目:用一个栈实现另一个栈的排序
	//sortStackByStack.go

	s := NewStack()
	s.Push(1)
	s.Push(3)
	s.Push(5)
	s.Push(2)
	s.Push(4)
	s.readAllStack()

	s.stackSort()
	s.readAllStack()

}

func test5()  {
	//题目:猫狗队列
	//catAndDogQueue.go
	s := NewCatDogQueue()
	v1 := NewDog("dog1")
	v2 := NewDog("dog2")
	v3 := NewDog("dog3")
	v4 := NewCat("cat1")
	v5 := NewCat("cat2")
	v6 := NewCat("cat3")


	fmt.Println(s.Add(v1))
	fmt.Println(s.Add(v4))
	fmt.Println(s.Add(v2))
	fmt.Println(s.Add(v5))
	fmt.Println(s.Add(v3))
	fmt.Println(s.Add(v6))
	fmt.Println(s.IsEmpty())

	//a := s.PollAll()
	//a:= s.PollCat()
	a:= s.PollDog()
	fmt.Println("a:",a)

	for _,v := range a {
		fmt.Println(v)
	}

}

func test4()  {
	//题目:如何仅用递归函数和栈操作逆序一个栈
	//reverseOrderStack.go
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(6)
	s.Push(5)
	s.Push(4)
	s.readAllStack()
	s.reverseOrder()

	s.readAllStack()
}

func test3()  {
	//twoStackToList.go
	//题目:由两个栈组成的队列
	sl := NewStackToList()

	sl.Push(1)
	sl.Push(2)
	sl.Push(3)
	sl.Push(4)


	fmt.Println(sl.Pop())
	fmt.Println(sl.Pop())
	fmt.Println(sl.Pop())
	fmt.Println(sl.Pop())
}

func test2()  {
	//getMainInn.go
	//题目:设计一个有getMin功能的栈

	gs := NewGetMinStack()
	gs.Push(22)
	gs.Push(2)
	gs.Push(33)
	gs.Push(4)
	gs.Push(54)
	fmt.Println(gs.GetMin())
}

func test1()  {
	//goStack.go
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}



