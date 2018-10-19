package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	h := head
	for h.Next != nil {
		h.Next, _ = reverseList(h.Next)
		h = h.Next
	}
}

/**
 * 逆转链表
 */
func reverseList(head *ListNode) (*ListNode, *ListNode) {
	if head.Next != nil {
		h, t := reverseList(head.Next)
		t.Next = head
		head.Next = nil
		return h, head
	}
	return head, head
}

func main() {
	l0 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}
	reorderList(l0)

	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	reorderList(l1)

	i := l0
	for i != nil {
		fmt.Println(i.Val)
		i = i.Next
	}

	i = l1
	for i != nil {
		fmt.Println(i.Val)
		i = i.Next
	}
}
