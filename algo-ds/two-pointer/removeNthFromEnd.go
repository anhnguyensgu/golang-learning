package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow := &ListNode{0, head}
	//using dummy node inorder to get pre-successor
	ans := slow
	fast := head

	count := n
	for count != 0 {
		fast = fast.Next
		count--
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return ans.Next
}

func main() {
	root := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	root = removeNthFromEnd(root, 5)

	for root != nil {
		fmt.Printf("%d ", root.Val)
		root = root.Next
	}
	fmt.Println()
}
