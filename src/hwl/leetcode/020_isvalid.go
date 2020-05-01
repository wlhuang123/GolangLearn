package leetcode

import "hwl/leetcode/check"

// Test020 判断字符串是否合法
/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
*/
func Test020() {
	Check020(isValid("{}"), true)
	Check020(isValid("[]"), true)
	Check020(isValid("{]"), false)
	Check020(isValid("]{"), false)
	Check020(isValid("]["), false)
	Check020(isValid("{[]}"), true)
	Check020(isValid("{{[]}}"), true)
	Check020(isValid("{}[[[]]]"), true)
	Check020(isValid("{}[]()"), true)
}

// Check020 .
func Check020(left, right bool) {
	check.IsBoolEqual(left, right)
}

func isValid(s string) bool {
	var array []string

	for i := 0; i < len(s); i++ {
		if len(array) == 0 {
			array = append(array, string(s[i]))
			continue
		}

		if isMatch(array[len(array)-1], string(s[i])) {
			array = array[0 : len(array)-1]
			continue
		}

		array = append(array, string(s[i]))
	}
	return len(array) == 0
}

// 判断左右是否匹配 比如{}属于匹配
func isMatch(left, right string) bool {
	if left == "{" && right == "}" {
		return true
	}

	if left == "[" && right == "]" {
		return true
	}

	if left == "(" && right == ")" {
		return true
	}

	return false
}
