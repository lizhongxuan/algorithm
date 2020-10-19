//生成窗口最大值数组
package main


//     ------->
//    push     pop
//    popFirst  pushLast

func getMaxWindow(arr []int,w int)[]int{
	if arr == nil || w<1||len(arr)<w {
		return nil
	}
	qmax := NewQueue()

	res := make([]int,len(arr)-w+1)
	index := 0

	for i:=0;i<len(arr) ;i++  {

		for !qmax.isEmpty() && arr[qmax.peekLast().(int)] <= arr[i] {
			qmax.Pop()
		}
		qmax.PushLast(i)
		if qmax.peekFirst().(int) == i-w {
			qmax.PopFirst()
		}

		if i >= w-1 {
			res[index] = arr[qmax.peekFirst().(int)]
			index += 1
		}
	}

	return res
}
