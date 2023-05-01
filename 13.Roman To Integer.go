package main

import "fmt"

func RomanToInt(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		value := charToInt(rune(s[i]))

		if i+1 < len(s) {
			nextValue := charToInt(rune(s[i+1]))
			if value < nextValue {
				result -= value
			} else {
				result += value
			}
		} else {
			result += value
		}
	}

	return result
}

func charToInt(c rune) int {
	switch c {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return -1
	}
}

func RomanToInt2(s string) int {
	var romanMap = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	var result = romanMap[s[len(s)-1]]

	fmt.Printf("%v\n", romanMap)
	for i := len(s) - 2; i >= 0; i-- {
		if romanMap[s[i]] < romanMap[s[i+1]] {
			result -= romanMap[s[i]]
		} else {
			result += romanMap[s[i]]
		}
	}
	return result
}
