package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	ans := &ListNode{Next: nil}
	for head != nil {
		temp := head
		head = head.Next

		temp.Next = ans.Next
		ans.Next = temp

	}

	return ans.Next
}

// a->b->c

// ans -> a
// ans -> next -> b -> c
