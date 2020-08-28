package leetcode

import (
	"hwl/leetcode/check"
	"strings"
)

/*
给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1
*/

// Test028 .
func Test028() {
	Check028("hello", "ll", 2)
	Check028("aaaaa", "bba", -1)
	Check028("adsjfl", "", 0)
}

// Check028 .
func Check028(haystack string, needle string, result int) {
	check.IsIntEqual(strStr(haystack, needle), result)
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	lenStack := len(haystack)
	for i := 0; i < lenStack; i++ {
		if strings.HasPrefix(haystack, needle) {
			return i
		}
		haystack = haystack[1:]
	}

	return -1
}
