//X题目:构造数组的MaxTree
package main

type binaryNode struct {
	value int
	left *binaryNode
	right *binaryNode
}

func NewBinaryNode(value int)*binaryNode  {
	return &binaryNode{
		value:value,
	}
}

func popStackSetMap(stack *Stack,m map[*binaryNode]*binaryNode)  {
	popNode := stack.Pop().(*binaryNode)
	if stack.isEmpty() {
		m[popNode]=nil
	}else {
		m[popNode] = stack.top.value.(*binaryNode)
	}
}

func getMaxTree(arr []int) *binaryNode {
	nArr := make([]*binaryNode,len(arr))
	for i:=0;i!=len(arr) ;i++  {
		nArr[i] = NewBinaryNode(arr[i])
	}

	stack := NewStack()
	lBigMap := make(map[*binaryNode]*binaryNode)
	//rBigMap := make(map[*binaryNode]*binaryNode)
	for i:=0;i!= len(nArr) ;i++  {
		curNode := nArr[i]
		for !stack.isEmpty() && stack.top.value.(*binaryNode).value < curNode.value  {
			popStackSetMap(stack,lBigMap)
		}
		stack.Push(curNode)
	}

	for !stack.isEmpty()  {
		popStackSetMap(stack,lBigMap)
	}

	for i:=len(nArr) - 1;i!=-1 ;i--  {
		//curNode := nArr[i]
	}

}