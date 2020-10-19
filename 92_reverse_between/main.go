package main

import "fmt"

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m>=n {
		return head
	}
	left := &ListNode{
		Next:head,
	}
	right := &ListNode{
		Next:head,
	}

	for i:=0;i<(m-1) ;i++  {
		left = left.Next
	}
	for i:=0;i<(n+1) ;i++  {
		right = right.Next
	}

	head2 := left.Next
	var cur *ListNode
	var tmp *ListNode
	var end *ListNode
	for head2 != nil && head2 != right {
		cur = head2
		head2 = head2.Next
		cur.Next = tmp
		tmp = cur
		if end == nil {
			end = cur
		}
	}
	left.Next = tmp
	end.Next = right

	if m==1 {
		return tmp
	}
	return head
}



func main() {
	list := createList()
	printList(list)
	list = reverseBetween(list,1,2)

	printList(list)
}

type ListNode struct {
	Val  int
	Next *ListNode
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