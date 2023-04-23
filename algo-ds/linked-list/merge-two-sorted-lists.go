package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var a = &ListNode{0, nil}
	var head *ListNode

	for list1 != nil || list2 != nil {
		if list1 == nil {
			a.Next = list2
			list2 = list2.Next
		} else if list2 == nil {
			a.Next = list1
			list1 = list1.Next
		} else if list2.Val < list1.Val {
			a.Next = list2
			list2 = list2.Next
		} else {
			a.Next = list1
			list1 = list1.Next
		}

		a = a.Next

	}

	return head.Next
}

func main() {

}
