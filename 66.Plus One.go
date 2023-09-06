package main

func PlusOne(digits []int) []int {
	carry := 1 // Start with a carry of 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += carry
		carry = digits[i] / 10 // Get the carryover for the next digit
		digits[i] %= 10        // Keep only the last digit
	}
	if carry > 0 { // If there is a carryover after processing all digits
		digits = append([]int{carry}, digits...) // Add a new digit with the carryover at the beginning
	}
	return digits
}

func PlusOne2(digits []int) []int {
	var n int = len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	var a = make([]int, n+1)
	a[0] = 1
	return a

}

func PlusOne3(digits []int) []int {

	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0

	}
	digits = append([]int{1}, digits...)
	return digits
}

func PlusOne4(digits []int) []int {
	var res int = 1
	for i := len(digits) - 1; i >= 0; i-- {
		s := digits[i]
		if s+res <= 9 {
			digits[i] += res
			break
		}
		digits[i] = 0
		if i == 0 {
			digits = append([]int{1}, digits...)
		}
	}
	return digits
}
