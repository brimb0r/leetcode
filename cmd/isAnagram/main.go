package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var alaphaCounter [26]int

	for i := 0; i < len(s); i++ {
		alaphaCounter[s[i]-'a']++
		alaphaCounter[t[i]-'a']--
	}
	for j := 0; j < len(alaphaCounter); j++ {
		if alaphaCounter[j] != 0 {
			return false
		}
	}
	return true
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	group := groupAnagrams(strs)
	fmt.Print(group)
}

type Key [26]byte

func strKey(str string) (key Key) {
	for i := range str {
		key[str[i]-'a']++
	}
	return
}

func groupAnagrams(strs []string) [][]string {
	groups := make(map[Key][]string)

	for _, v := range strs {
		key := strKey(v)
		groups[key] = append(groups[key], v)
	}

	result := make([][]string, 0, len(groups))
	for _, v := range groups {
		result = append(result, v)
	}
	return result
}
