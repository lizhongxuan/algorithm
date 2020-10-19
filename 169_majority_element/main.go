package main

import "fmt"

func majorityElement(nums []int) int {


	fmt.Println("---",nums)
	if (len(nums)<=2 && nums[0] == nums[len(nums)-1]) {
		return nums[0];
	}

	mid := len(nums)/2
	fmt.Println()
	fmt.Println(nums)
	fmt.Println("mid:",mid)
	fmt.Println()
	fmt.Println("len(nums)-1:",len(nums)-1)
	leftNum := majorityElement(nums[0:mid])
	rightNum := majorityElement(nums[mid:])
	fmt.Println("----")

	if leftNum == rightNum {
		return leftNum
	}
	leftCount := countNum(nums,leftNum)
	rightCount := countNum(nums,rightNum)

	if leftCount >= rightCount {
		return leftNum
	}
	return rightNum

}

func countNum(nums []int , num int) int {
	count := 0
	for i:=0;i<len(nums) ;i++  {
		if nums[i] == num {
			count++
		}
	}
	return count
}


func main() {
	fmt.Println(majorityElement([]int{10,9,9,9,10}))
}
