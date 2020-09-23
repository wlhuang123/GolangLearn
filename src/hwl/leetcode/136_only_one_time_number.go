package leetcode

import (
	"hwl/leetcode/check"
)

/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:
输入: [2,2,1]
输出: 1

示例 2:
输入: [4,1,2,1,2]
输出: 4
*/

// Test136 .
func Test136() {
	Check136([]int{2, 2, 1}, 1)
	Check136([]int{4, 1, 2, 1, 2}, 4)
}

// Check136 .
func Check136(nums []int, result int) {
	check.IsIntEqual(singleNumber(nums), result)
}

// 利用异或逻辑。两个相同的数字异或结果是0 。 0对一个数的异或是数字的本身
// 所以所有数字异或一下，相同的数字互相抵消。最后剩下期望结果和0的异或，就得到了期望结果
func singleNumber(nums []int) int {
	var result int

	for _, v := range nums {
		result = result ^ v
	}

	return result
}
