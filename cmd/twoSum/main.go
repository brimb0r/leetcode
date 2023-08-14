package main

import (
	"fmt"
	"sort"
	"time"
)

// brute force
func twoSumBrute(nums []int, target int) []int {
	n := len(nums) // lenght of slice
	for i := 0; i < n-1; /* reduce by one to */ i++ {
		for j := i + 1; /* from start add next */ j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	// if not found return empty slice
	return []int{}
}

// hashmap

func twoSum(nums []int, target int) []int {
	// make a map length of nums
	indexMap := make(map[int]int, len(nums))
	// loop over slice where cidx = index is slice and curr num is the value
	for currIndex, currNum := range nums {
		// here we look to see if the target - the current number is in the map e.g
		// 9-7, if that number is found return the index of where it was found and the current index
		if reqIdx, isPresent := indexMap[target-currNum]; isPresent {
			return []int{reqIdx, currIndex}
		}
		indexMap[currNum] = currIndex
	}
	return []int{}
}

func threeSumBrute(nums []int) [][]int {
	/*
		take nums[0] + nums[1] + nums[2] = 0 otherwise the trip does not work
	*/

	// lenght of our slice
	n := len(nums)

	// init a multidem map of ints
	tripMap := make(map[[3]int][]int)

	// this is brute force, we need 3 loops for add up

	// i take all but last two in slice
	for i := 0; i < n-2; i++ {
		// j take all but last 1
		for j := i + 1; j < n-1; j++ {
			// k take all
			for k := j + 1; k < n; k++ {
				// our triplet here is going to be the three from the slice and sorted to find only unq
				trip := [3]int{nums[i], nums[j], nums[k]}
				sort.Ints(trip[:])
				// add them up and see if it == 0
				if nums[i]+nums[j]+nums[k] == 0 {
					// if true append into our map
					tripMap[trip] = []int{nums[i], nums[j], nums[k]}
				}
			}
		}
	}
	var results [][]int
	for _, triplet := range tripMap {
		results = append(results, triplet)
	}
	return results
}

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	var results [][]int
	resultsChan := make(chan []int, n)

	for numberIndex := range nums {
		// skip dups from left
		if numberIndex > 0 && nums[numberIndex] == nums[numberIndex-1] {
			continue
		}
		number2Idx := numberIndex + 1
		number3Idx := n - 1
		go routineCheck(nums, number2Idx, number3Idx, numberIndex, resultsChan)
	}
	for result := range resultsChan {
		results = append(results, result)
	}

	return results
}

func routineCheck(nums []int, number2Idx int, number3Idx int, numberIndex int, out chan<- []int) {
	for number2Idx < number3Idx {
		sum := nums[numberIndex] + nums[number2Idx] + nums[number3Idx]
		if sum == 0 {
			out <- []int{nums[numberIndex], nums[number2Idx], nums[number3Idx]}
			number3Idx--
			// Skip all duplicates from right
			for number2Idx < number3Idx && nums[number3Idx] == nums[number3Idx+1] {
				number3Idx--
			}
		} else if sum > 0 {
			// Decrement num3Idx to reduce sum value
			number3Idx--
		} else {
			// Increment num2Idx to increase sum value
			number2Idx++
		}
	}
	close(out)
}

func indexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func main() {
	//twoSumCheck := twoSum([]int{2, 7, 11, 15}, 9)
	start := time.Now()
	threeSumCheck := threeSum([]int{0, 0, 0, 0})
	fmt.Printf("Output %d - duration  %v\n", threeSumCheck, time.Since(start))
}
