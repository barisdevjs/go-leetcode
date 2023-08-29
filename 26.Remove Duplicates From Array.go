package main

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}

func RemoveDuplicates2(nums []int) int {
	if len(nums) <= 1 {
		return 1
	}

	slow := 0
	fast := 1
	n := len(nums)

	for fast < n {
		if nums[fast] == nums[slow] {
			fast++
		} else {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}

func RemoveDuplicates3(nums []int) int {
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[j] == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		} else {
			j++
		}
	}
	return len(nums)
}
