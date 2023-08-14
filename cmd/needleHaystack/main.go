package main

import (
	"fmt"
	"strings"
)

type NeedleHaystack struct {
	haystack string
	needle   string
}

func main() {
	initNeddleHaysack := &NeedleHaystack{
		haystack: "sadbutsad",
		needle:   "sad",
	}
	findNeedle := initNeddleHaysack.strStr()
	if findNeedle < 0 {
		fmt.Printf("Needle was not found %d", findNeedle)
	} else {
		fmt.Printf("Needle found at index %d", findNeedle)
	}
}

func (nh *NeedleHaystack) strStr() int {
	// base case = is needle even in haystack?
	if !strings.Contains(nh.haystack, nh.needle) {
		return -1
	} else {
		foundAt := strings.Index(nh.haystack, nh.needle)
		return foundAt
	}
}
