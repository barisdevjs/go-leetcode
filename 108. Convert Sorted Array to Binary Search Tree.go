package main

func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = SortedArrayToBST(nums[:mid])
	root.Right = SortedArrayToBST(nums[mid+1:])
	return root
}

func SortedArrayToBST2(nums []int) *TreeNode {
	if nums == nil {
		return nil
	}
	return helperSortedArrayToBST(nums, 0, len(nums)-1)
}

func helperSortedArrayToBST(nums []int, low int, high int) *TreeNode {
	if low > high { // Done
		return nil
	}
	mid := low + (high-low)/2
	node := new(TreeNode)
	node.Val = nums[mid]
	node.Left = helperSortedArrayToBST(nums, low, mid-1)
	node.Right = helperSortedArrayToBST(nums, mid+1, high)
	return node
}
