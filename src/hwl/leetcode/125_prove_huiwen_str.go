package leetcode

import (
	"hwl/leetcode/check"
	"unicode"
)

/*
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:
输入: "A man, a plan, a canal: Panama"
输出: true

示例 2:
输入: "race a car"
输出: false
*/

// Test125 .
func Test125() {
	Check125("A man, a plan, a canal: Panama", true)
	Check125("ace a car", false)
	Check125("", true)
}

// Check125 .
func Check125(s string, result bool) {
	check.IsBoolEqual(isPalindrome(s), result)
}

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	begin := 0
	end := len(s) - 1

	for {
		if begin >= end {
			return true
		}

		if !unicode.IsDigit(rune(s[begin])) && !unicode.IsLetter(rune(s[begin])) {
			begin++
			continue
		}

		if !unicode.IsDigit(rune(s[end])) && !unicode.IsLetter(rune(s[end])) {
			end--
			continue
		}

		if unicode.IsDigit(rune(s[begin])) && unicode.IsDigit(rune(s[end])) {
			if s[begin] == s[end] {
				begin++
				end--
				continue
			}
			return false
		}

		if unicode.IsLetter(rune(s[begin])) && unicode.IsLetter(rune(s[end])) {
			if unicode.ToLower(rune(s[begin])) == unicode.ToLower(rune(s[end])) {
				begin++
				end--
				continue
			}
			return false
		}
		return false
	}
}
