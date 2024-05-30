package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}
	reverseList(head)
	printList(head)
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Printf("nil")
}

/**
 * Definition for singly-linked list.

 */
func reverseList(head *ListNode) (prev *ListNode) {
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return
}
