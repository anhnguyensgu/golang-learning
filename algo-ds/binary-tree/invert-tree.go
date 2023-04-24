package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	var temp *TreeNode
	if root.Left != nil {
		temp = invertTree(root.Left)
	}

	if root.Right != nil {
		invertTree(root.Right)
	}

	root.Left = root.Right
	root.Right = temp

	return root
}

func main() {

}
