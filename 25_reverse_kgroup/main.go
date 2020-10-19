package main

import "fmt"

func main() {

	list := createList()

	list,list2 := onceReverseKGroup(list,4)
	printList(list)
	fmt.Println("---")
	printList(list2)

}

type ListNode struct {
	     Val int
	     Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	len := lenList(head)
	l := &ListNode{
	}
	t := l
	head2:= head
	for len > k {
		head2,ol := onceReverseKGroup(head2,k)
		l.Next = ol
		len -= k
	}
	if len != 0 {
		onceReverseKGroup(head,len)
	}


	return t.Next
}

func onceReverseKGroup(head *ListNode, k int) (*ListNode,*ListNode) {
	left := &ListNode{
	}

	for i:=0;i<k ;i++  {
		if head == nil {
			return nil,nil
		}

		tmp := &ListNode{
			Next:head,
		}

		head = head.Next
		tmp.Next.Next = left
		left = tmp.Next
	}
	return head,left.Next
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
		fmt.Println(head.Val)
		head = head.Next
	}
}

func lenList(head *ListNode)int  {
	len :=0
	for head != nil {
		len++
		head = head.Next
	}

	return len
}