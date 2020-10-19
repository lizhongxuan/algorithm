package main

//func maxRecSize(m [][]int)int {
//	if m == nil || len(m)==0 || len(m[0])==0 {
//		return 0
//	}
//	maxArea := 0
//	height := make([]int,len(m[0]))
//	for i:=0;i<len(m) ;i++  {
//		for j:=0;j<len(m[0]);j++  {
//			if m[i][j] == 0 {
//				height[j] = 0
//			}else {
//				height[j] = height[j]+1
//			}
//		}
//		maxArea = math.Max()
//	}
//	return maxArea
//}
//
//func maxRecFromBttom(height []int)  {
//	if height == nil || len(height) == 0 {
//		return 0
//	}
//	maxArea := 0
//	stack := NewStack()
//	for i:=0;i<len(height) ;i++  {
//		for !stack.isEmpty() && height[i] <= height[stack.top.value.(int)]  {
//			j := stack.Pop().(int)
//		}
//	}
//}
