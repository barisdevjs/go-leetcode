package main

import "strconv"

func IsPalindrome(x int) bool {
	str := strconv.Itoa(x)
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	return str == reversedStr
}

func IsPalindrome2(x int) bool {
	int_string := strconv.Itoa(x)
	n := len(int_string)
	for i := 0; i < n/2; i++ {
		if rune(int_string[i]) != rune(int_string[n-1-i]) {
			return false
		}
	}
	return true
}

func IsPalindrome3(x int) bool {
	if x < 0 {
		return false
	} else if x <= 9 {
		return true
	} else if x%10 == 0 {
		return false
	}

	var y int
	for x > y {
		r := x % 10
		x = x / 10
		y = y*10 + r

		if x == y || x/10 == y {
			return true
		}
	}
	return false
}

func IsPalindrome4(x int) bool {
	if x < 0 {
		return false
	}
	r := 0
	for z := x; z > 0; z /= 10 {
		tmp := z % 10
		r *= 10
		r += tmp

	}
	return r == x
}
