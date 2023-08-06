package validatebinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const MAX_VALUE int = 3 * 10e9
const MIN_VALUE int = -3 * 10e9

func isValidBST(root *TreeNode) bool {
	return dfs(root, MAX_VALUE, MIN_VALUE)
}

func dfs(root *TreeNode, maxValue, minValue int) bool {
	if root == nil {
		return true
	}
	if root.Val >= maxValue || root.Val <= minValue {
		return false
	}

	leftMax := dfs(root.Left, root.Val, minValue)
	rightMax := dfs(root.Right, maxValue, root.Val)
	return leftMax && rightMax
}
