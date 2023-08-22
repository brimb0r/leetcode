package main

import "fmt"

func main() {
	romanToIntConversion := romanToInt("MCMXCIV")
	fmt.Print(romanToIntConversion, "\n")
	intToRomaneConversion := intToRoman(1994)
	fmt.Print(intToRomaneConversion, "\n")
}

/*
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
*/

/*
I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
*/
var romanToIntMap = map[uint8]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {

	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return romanToIntMap[s[0]]
	}

	var value, lastvalue int
	for i := len(s) - 1; i >= 0; i-- {
		if romanToIntMap[s[i]] < lastvalue {
			value -= romanToIntMap[s[i]]
		} else {
			value += romanToIntMap[s[i]]
		}
		lastvalue = romanToIntMap[s[i]]
	}
	return value
}

func intToRoman(num int) string {
	var out string
	inputs := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	for _, v := range inputs {
		if num/v > 0 {
			roman := intToRomanMap[v]
			for k := 0; k < num/v; k++ {
				out += roman
			}
			num = num % v
		}
	}
	return out
}

var intToRomanMap = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
	900:  "CM",
	400:  "CD",
	90:   "XC",
	40:   "XL",
	9:    "IX",
	4:    "IV",
}
