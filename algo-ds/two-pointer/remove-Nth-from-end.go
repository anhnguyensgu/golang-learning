package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return *slow
}

func main() {
	child2 := ListNode{3, &ListNode{4, &ListNode{5, nil}}}
	child := ListNode{2, &child2}
	root := ListNode{1, &child}

	fmt.Println(middleNode(&root).Val)

}
