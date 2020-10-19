package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	cur := &ListNode{}
	var tmp *ListNode
	for head != nil  {
		cur = head
		head = head.Next
		cur.Next = tmp
		tmp = cur
	}
	return tmp
}

func main() {
	list := createList()
	printList(list)
	list = reverseList(list)

	printList(list)

}

func createList()*ListNode  {
	list := &ListNode{
		Val:1,
	}
	n := list

	for i:=2; i<9; i++ {
		n.Next = &ListNode{
			Val:i,
		}
		n = n.Next
	}
	return list
}

func printList(head *ListNode)  {
	for head != nil {
		fmt.Printf("->%d",head.Val)
		head = head.Next
	}
	fmt.Println("---<<<")
}