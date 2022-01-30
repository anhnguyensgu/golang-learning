package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
var ans = 0

func abs(a int, b int) int {
	if a-b > 0 {
		return a - b
	} else {
		return b - a
	}
}

func maxAncestorDiff(root *TreeNode) int {
	ans = 0
	helper(root)
	return ans
}

func minValue(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maxValue(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func helper(root *TreeNode) (int, int) {
	minLeft, minRight := root.Val, root.Val
	maxRight, maxLeft := root.Val, root.Val
	if root.Left != nil {
		minLeft, maxLeft = helper(root.Left)
	}

	if root.Right != nil {
		minRight, maxRight = helper(root.Right)
	}

	diffLeft, diffRight := abs(root.Val, minLeft), abs(root.Val, minRight)

	ans = maxValue(ans, maxValue(diffRight, diffLeft))

	diffLeft, diffRight = abs(root.Val, maxLeft), abs(root.Val, maxRight)

	ans = maxValue(ans, maxValue(diffRight, diffLeft))

	return minValue(root.Val, minValue(minLeft, minRight)), maxValue(root.Val, maxValue(maxRight, maxLeft))
}

func main() {
	root := &TreeNode{2,
		nil, &TreeNode{0,
			nil, &TreeNode{4,
				nil, &TreeNode{3, nil, &TreeNode{1, nil, nil}}}}}

	maxAncestorDiff(root)
	fmt.Println(ans)
}
