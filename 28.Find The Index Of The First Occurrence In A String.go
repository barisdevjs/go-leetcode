package main

import (
	"strings"
)

func StrStr(haystack string, needle string) int {
	hLen := len(haystack)
	needLen := len(needle)

	if hLen < needLen {
		return -1
	}

	for i := 0; i < hLen; i++ {
		temp := haystack[i:]
		if strings.HasPrefix(temp, needle) {
			return i
		}
	}

	return -1
}

func StrStr2(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	h := len(haystack)
	n := len(needle)
	for i := 0; i < h-n+1; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}

func StrStr3(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
