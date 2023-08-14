package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	// base case of zero is just zero
	if x == 0 {
		return 0
	}

	// remove  zeros at the end of number
	for {
		if x%10 == 0 {
			x = x / 10
		} else {
			break
		}
	}
	// convert to string
	asString := strconv.Itoa(x)
	reverseString := make([]byte, len(asString))
	end := 0
	if asString[0] == '-' {
		end = 1
	}
	j := 0
	if end == 1 {
		j = 1
		reverseString[0] = '-'
	}

	// reverse by swap bytes
	for i := len(asString) - 1; i >= end; i-- {
		reverseString[j] = asString[i]
		j++
	}
	// convert back to int
	asInt, _ := strconv.Atoi(string(reverseString))

	if asInt < math.MinInt32 || asInt > math.MaxInt32 {
		return 0
	}
	return asInt
}

func main() {
	c := reverse(-3243564634534644)
	fmt.Print(c)
}

func isPal(x int) bool {
	var re int
	tmp := x
	for tmp > 0 {
		re = re*10 + tmp%10
		tmp = tmp / 10
	}
	return x == re
}

func reverse1(x int) int {
	result, sign := 0, 1

	if x < 0 {
		sign = -1
		x = -x
	}

	for x > 0 {
		remainder := x % 10
		result = result*10 + remainder
		x = x / 10
	}

	var checkInt int = int(int32(result))

	if checkInt != result {
		return 0
	}

	return result * sign

}
