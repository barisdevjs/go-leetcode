package main

import "fmt"

func TwoSum(nums []int, target int) []int {
	result := make(map[int]int, 2)

	for i, num := range nums {
		compliment := target - num
		if j, ok := result[compliment]; ok {
			return []int{i, j}
		}
		fmt.Println(result)
		result[num] = i
	}
	return nil
}

func TwoSum2(nums []int, target int) []int {
	// O(n)
	for i := 0; i < len(nums)-1; i++ {
		// O(n)
		for j := i + 1; j < len(nums); j++ {
			// Time: O(1)
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
