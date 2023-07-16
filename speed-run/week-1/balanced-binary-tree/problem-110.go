package balancedbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var b bool

func isBalanced(root *TreeNode) bool {
	b = true
	if root == nil {
		return b
	}
	dfs(root)
	return b
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 1
	}
	leftLv := dfs(root.Left)
	rightLv := dfs(root.Right)
	diff := leftLv - rightLv
	if diff < 0 {
		diff *= -1
	}
	if diff > 1 {
		b = false
	}
	if leftLv > rightLv {
		return leftLv + 1
	}
	return rightLv + 1
}
