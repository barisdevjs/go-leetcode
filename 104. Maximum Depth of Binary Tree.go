package main

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := MaxDepth(root.Left)
	rightDepth := MaxDepth(root.Right)
	return 1 + max(leftDepth, rightDepth)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
