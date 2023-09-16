package main

import (
	"fmt"
	"sort"
)

func Merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	copy(nums1[m:], nums2)
	// nums1 = append(nums1[:m], nums2...)
	sort.Ints(nums1)
}

func Merge2(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	i, j := 0, 0

	for i < m && j < n {
		if nums2[j] < nums1[i] {
			copy(nums1[i+1:], nums1[i:m])
			fmt.Println("XX", nums1)
			nums1[i] = nums2[j]
			j++
			m++
		} else {
			i++
		}
	}
	if j < n {
		nums1 = append(nums1[:m], nums2[j:]...)
	}
	fmt.Println(nums1)
}

func Merge3(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for ; j >= 0; k-- {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
}
