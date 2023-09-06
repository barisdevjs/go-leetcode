package main

import (
	"fmt"
	"strings"
)

func LengthOfLastWord(s string) int {

	s = strings.TrimSpace(s)
	res := 0

	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != string(' ') {
			res++
		} else {
			break
		}
	}
	return res
}

func LengthOfLastWord2(s string) int {
	s = strings.TrimRight(s, " ")
	fmt.Println(s)
	return len(s) - 1 - strings.LastIndex(s, " ")
}

func LengthOfLastWord3(s string) int {
	str := strings.Fields(s) //tokenize
	fmt.Println(str[len(str)-1])
	return len(str[len(str)-1])
}
