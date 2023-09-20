package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
	var nums = []int{}
	if root == nil {
		return nums
	}

	nums = append(nums, InorderTraversal(root.Left)...)
	nums = append(nums, root.Val)
	nums = append(nums, InorderTraversal(root.Right)...)
	return nums
}

func InorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0, 101)
	if root != nil {
		res = help(root, res)
	}

	return res
}

func help(root *TreeNode, res []int) []int {
	if root.Left != nil {
		res = help(root.Left, res)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		res = help(root.Right, res)
	}
	return res
}
