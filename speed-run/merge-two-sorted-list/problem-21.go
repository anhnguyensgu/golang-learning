package mergetwosortedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l *ListNode, r *ListNode) *ListNode {
	tail := &ListNode{Val: -1}
	ans := tail

	for l != nil || r != nil {
		if l == nil {
			tail.Next = r
			r = r.Next
		} else if r == nil {
			tail.Next = l
			l = l.Next
		} else {
			if l.Val < r.Val {
				tail.Next = l
				l = l.Next
			} else {
				tail.Next = r
				r = r.Next
			}
		}
		tail = tail.Next
	}

	return ans.Next
}
