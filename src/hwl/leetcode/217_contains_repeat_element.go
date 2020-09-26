package leetcode

import (
	"hwl/leetcode/check"
)

/*
217. 存在重复元素
如果任意一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。

示例 1
输入: [1,2,3,1]
输出: true

示例 2:
输入: [1,2,3,4]
输出: false

示例 3:
输入: [1,1,1,3,3,4,3,2,4,2]
输出: true
*/

// Test217 .
func Test217() {
	Check217([]int{1, 2, 3, 1}, true)
	Check217([]int{1, 2, 3, 4}, false)
	Check217([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true)
}

// Check217 .
func Check217(nums []int, expectResult bool) {
	check.IsBoolEqual(containsDuplicate(nums), expectResult)
}

func containsDuplicate(nums []int) bool {
	tempMap := make(map[int]int)

	for _, num := range nums {
		_, ok := tempMap[num]
		if ok {
			return true
		}
		tempMap[num] = 0
	}
	return false
}
