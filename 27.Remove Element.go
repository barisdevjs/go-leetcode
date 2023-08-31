package main

import "fmt"

func RemoveElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	j := 0
	for _, num := range nums {
		if num != val {
			nums[j] = num
			j++
		}
	}

	return j
}

func RemoveElement2(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	newNums := make([]int, 0, len(nums))
	for _, num := range nums {
		if num != val {
			newNums = append(newNums, num)
		}
	}

	copy(nums, newNums)

	return len(newNums)
}

func RemoveElement3(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if val == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
			fmt.Println(nums)
		}
	}
	return len(nums)
}

func RemoveElement4(nums []int, val int) int {
	i := 0
	for _, v := range nums {
		if v != val {
			nums[i] = v
			i++
		}
	}
	return i
}
