package leetcode

import (
	"hwl/leetcode/check"
)

/*
219. 存在重复元素 II
给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的 绝对值 至多为 k。

示例 1:
输入: nums = [1,2,3,1], k = 3
输出: true

示例 2:
输入: nums = [1,0,1,1], k = 1
输出: true

示例 3:
输入: nums = [1,2,3,1,2,3], k = 2
输出: false
*/

// Test219 .
func Test219() {
	Check219([]int{1, 2, 3, 1}, 3, true)
	Check219([]int{1, 0, 1, 1}, 1, true)
	Check219([]int{1, 2, 3, 1, 2, 3}, 2, false)
}

// Check219 .
func Check219(nums []int, k int, expectResult bool) {
	check.IsBoolEqual(containsNearbyDuplicate(nums, k), expectResult)
}

func containsNearbyDuplicate(nums []int, k int) bool {
	tempMap := make(map[int]int) // key num  value: index
	for index, num := range nums {
		saveIndex, ok := tempMap[num]
		if ok {
			if index-saveIndex <= k {
				return true
			}
		}
		tempMap[num] = index
	}
	return false
}
