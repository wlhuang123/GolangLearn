package leetcode

import (
	"hwl/leetcode/check"
)

/*
给定一个整数，编写一个函数来判断它是否是 2 的幂次方。

示例 1:
输入: 1
输出: true
解释: 20 = 1

示例 2:
输入: 16
输出: true
解释: 24 = 16

示例 3:
输入: 218
输出: false
*/

// Test231 .
func Test231() {
	Check231(1, true)
	Check231(16, true)
	Check231(218, false)
}

// Check231 .
func Check231(n int, expectResult bool) {
	check.IsBoolEqual(isPowerOfTwo(n), expectResult)
}

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}

	var count int
	for n != 0 {
		if n&(0x1) == 1 {
			count++
		}
		n = n >> 1

	}
	return count == 1
}

func isPowerOfTwo2(n int) bool {
	if n <= 0 {
		return false
	}

	return n&(n-1) == 0
}
