package leetcode

import (
	"hwl/leetcode/check"
)

/*
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1：
输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。

示例 2：
输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。

示例 3：
输入：digits = [0]
输出：[1]
*/

// Test066 .
func Test066() {
	Check066([]int{1, 2, 3}, []int{1, 2, 4})
	Check066([]int{4, 3, 2, 1}, []int{4, 3, 2, 2})
	Check066([]int{0}, []int{1})
	Check066([]int{9}, []int{1, 0})
}

// Check066 .
func Check066(input, expect []int) {
	check.IsTowSliceTheSame(plusOne(input), expect)
}

// 技巧，无需每一位都做运算和取余，因为只有9才会发生进位
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			for j := len(digits) - 1; j > i; j-- {
				digits[j] = 0
			}
			return digits
		}
	}
	res := make([]int, len(digits)+1)
	res[0] = 1
	return res
}
