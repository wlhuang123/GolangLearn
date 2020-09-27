package leetcode

import (
	"hwl/leetcode/check"
)

// Test242 .
func Test242() {
	Check242("cat", "tac", true)
}

// Check242 .
func Check242(s, t string, expectResult bool) {
	check.IsBoolEqual(isAnagram(s, t), expectResult)
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letterMap := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		letterMap[s[i]]++
		letterMap[t[i]]--
	}

	for _, v := range letterMap {
		if v != 0 {
			return false
		}
	}
	return true
}
