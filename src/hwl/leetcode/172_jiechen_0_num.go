package leetcode

import (
	"hwl/leetcode/check"
)

/*
给定一个整数 n，返回 n! 结果尾数中零的数量。

示例 1:

输入: 3
输出: 0
解释: 3! = 6, 尾数中没有零。
示例 2:

输入: 5
输出: 1
解释: 5! = 120, 尾数中有 1 个零.
*/

// Test172 .
func Test172() {
	Check172(3, 0)
	Check172(5, 1)
	Check172(30, 7)
}

// Check172 .
func Check172(n, expectResult int) {
	check.IsIntEqual(trailingZeroes(n), expectResult)
}

// 原理 有5才能产生0 比如15=3*5 20=4*5 这些都有1个5
// 特殊的 25=5*5，25的倍数有两个5 125的倍数有3个5
func trailingZeroes(n int) int {
	count := 0
	for n > 0 {
		count += n / 5
		n = n / 5
	}
	return count
}
