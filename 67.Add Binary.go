package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func AddBinary(a string, b string) string {
	// Initialize variables for the result and the carry
	result := ""
	carry := 0

	// Pad the shorter string with zeros on the left to make the two strings the same length
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	a = strings.Repeat("0", maxLen-len(a)) + a
	b = strings.Repeat("0", maxLen-len(b)) + b

	fmt.Println(a, b)
	// Iterate over the digits from right to left
	for i := maxLen - 1; i >= 0; i-- {
		sum := int(a[i]-'0') + int(b[i]-'0') + carry
		if sum >= 2 {
			result = strconv.Itoa(sum-2) + result
			carry = 1
		} else {
			result = strconv.Itoa(sum) + result
			carry = 0
		}
	}

	// If there is a carry at the end, add it to the result
	if carry == 1 {
		result = "1" + result
	}

	// Return the final result
	return result
}

func AddBinary2(a string, b string) string {
	aInt, bInt, sum := new(big.Int), new(big.Int), new(big.Int)
	aInt.SetString(a, 2)
	bInt.SetString(b, 2)
	sum.Add(aInt, bInt)
	fmt.Println(aInt, bInt, sum)
	return sum.Text(2)
}
