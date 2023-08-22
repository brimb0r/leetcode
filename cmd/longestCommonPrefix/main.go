package main

import (
	"fmt"
)

func main() {
	longestCommon := longestCommonPrefix([]string{"dog", "racecar", "car"})
	fmt.Print(longestCommon, "\n")
}

func longestCommonPrefix(strs []string) string {
	// loop all index
	for i := 0; ; i++ {
		// each string in slice
		for _, str := range strs {
			// if i = len break, and if any elements dont equal break
			if i == len(str) || str[i] != strs[0][i] {
				// return first string and sub array up to i
				return strs[0][:i]
			}
		}
	}
}
