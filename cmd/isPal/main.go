package main

import (
	"fmt"
	"regexp"
	"strings"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func isPalindromeSlow(s string) bool {
	s = strings.ToLower(s)
	s = nonAlphanumericRegex.ReplaceAllString(s, " ")
	s = strings.ReplaceAll(s, " ", "")
	end := len(s) - 1

	fmt.Println(len(s))

	if len(s) == 0 {
		return true
	}

	if s[0] != s[end] {
		return false
	}

	for i, j := 0, end; i < end && j > 0; i, j = i+1, j-1 {
		fmt.Println(s[i], s[j])
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func main() {
	validate := isPalindromeSlow(".,")
	fmt.Println(validate)
}
