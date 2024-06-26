package main

func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)
	return abs(leftHeight-rightHeight) <= 1 && IsBalanced(root.Left) && IsBalanced(root.Right)
}

func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftHeight := getHeight(node.Left)
	rightHeight := getHeight(node.Right)
	return maxIsBalanced(leftHeight, rightHeight) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func maxIsBalanced(x, y int) int {
	if x > y {
		return x
	}
	return y
}
