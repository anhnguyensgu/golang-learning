package linkedlistcycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast := head.Next

	for fast != nil && fast.Next != nil && fast != head {
		head = head.Next
		fast = fast.Next.Next
	}

	return fast == head
}
