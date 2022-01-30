package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans1 *TreeNode
var ans2 *TreeNode
var prev *TreeNode

func dfs(node *TreeNode) {
	if node.Left != nil {
		dfs(node.Left)
	}
	if node.Val < prev.Val {
		if ans1 == nil {
			ans1 = prev
		}
		ans2 = node
	}
	prev = node

	if node.Right != nil {
		dfs(node.Right)
	}
}

func recoverTree(root *TreeNode) {
	prev = &TreeNode{-2147483648, nil, nil}
	ans1 = nil
	ans2 = nil
	dfs(root)
	ans1.Val, ans2.Val = ans2.Val, ans1.Val
}

func main() {
	root := &TreeNode{3,
		&TreeNode{1, nil, nil},
		&TreeNode{4, &TreeNode{2, nil, nil}, nil}}
	recoverTree(root)
}
