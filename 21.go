package main

import "fmt"

func main() {
	var l10 *ListNode = new(ListNode)
	l10.Val = 1
	l11 := new(ListNode)
	l11.Val = 3
	l10.Next = l11
	l12 := new(ListNode)
	l12.Val = 5
	l11.Next = l12

	var l20 *ListNode = new(ListNode)
	l20.Val = 2
	l21 := new(ListNode)
	l21.Val = 4
	l20.Next = l21
	l22 := new(ListNode)
	l22.Val = 6
	l21.Next = l22
	var list *ListNode = mergeTwoLists(l10,l20)

	fmt.Println(*list)
}


// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode = new(ListNode)
	var next *ListNode = head
	for nil != l1 && nil != l2 {
		if l1.Val  < l2.Val {
			next.Next = l1
			l1 = l1.Next
		}else{
			next.Next = l2
			l2 = l2.Next
		}
		next = next.Next
	}
	if nil != l1{
		next.Next = l1
	}else{
		next.Next = l2
	}
	return head.Next
}
