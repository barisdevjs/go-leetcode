package main

func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	}
	arr := make([]int, n)
	arr[0], arr[1] = 1, 2
	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n-1]
}

func ClimbStairs2(n int) int {
	a, b := 1, 1
	for ; n > 1; n-- {
		a, b = b, a+b
	}
	return b
}
