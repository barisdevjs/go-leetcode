package main

import (
	"sort"
)

func SearchInsert(nums []int, target int) int {

	if target > nums[len(nums)-1] {
		return len(nums)
	}
	if target < nums[0] {
		return 0
	}

	for i, v := range nums {
		if v == target {
			return i
		}
		if v > target {
			return i
		}

	}
	return -1
}

func SearchInsert2(nums []int, target int) int {
	return sort.SearchInts(nums, target)
}

func SearchInsert3(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := l + (r-l)/2
		v := nums[m]
		switch {
		case v < target:
			l = m + 1
		case v > target:
			r = m - 1
		default:
			return m
		}
	}
	return l
}

func SearchInsert4(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}
