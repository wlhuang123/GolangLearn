package leetcode

import "hwl/leetcode/check"

// Test009 给定一个数,判断是不是回文数
func Test009() {
	Check009(isPalindrome(-121), false)
	Check009(isPalindrome(0), true)
	Check009(isPalindrome(-1), false)
	Check009(isPalindrome(5), true)
	Check009(isPalindrome(50), false)
	Check009(isPalindrome(1221), true)
}

// Check009 .
func Check009(left, right bool) {
	check.IsBoolEqual(left, right)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	if x%10 == 0 {
		return false
	}

	var cur int

	for x > cur {
		cur = cur*10 + x%10
		x = x / 10
	}

	// x == cur/10 即去除中位数,当是奇数的时候
	return x == cur || x == cur/10
}
