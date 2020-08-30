package leetcode

import (
	"fmt"
	"hwl/tool"
	"hwl/tool/logs"
)

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

给定数组 nums = [-1, 0, 1, 2, -1, -4]，
满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

*/

// Test015 .
func Test015() {
	if tool.IsTowSliceEqual([]int{0, 0, 0}, []int{4, 0, -4}) {
		fmt.Println("equal")
	}

	Check015([]int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0})
	Check015([]int{0, 0, 0})
	Check015([]int{-1, 0, 1, 2, -1, -4})
	Check015([]int{0, 0, 0, 0})
}

// Check015 .
func Check015(nums []int) {
	result := threeSum(nums)
	logs.Println(result)
}

func threeSum(nums []int) [][]int {
	var result [][]int

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			target := -1 * (nums[i] + nums[j])
			if IsArrExist(nums[j+1:], target) {
				result = append(result, []int{nums[i], nums[j], target})
			}
		}
	}

	fmt.Println("begin result:", result)
	return RemoveRepeat(result)
}

// RemoveRepeat .
func RemoveRepeat(arr [][]int) [][]int {
	var newArr [][]int
	for i := 0; i < len(arr)-1; i++ {
		fmt.Print("index:", i, " ")
		if !IsTwoArrExist(arr[i+1:], arr[i]) {
			fmt.Println(" append:", arr[i])
			newArr = append(newArr, arr[i])
		}
	}
	newArr = append(newArr, arr[len(arr)-1])
	return newArr
}

// IsTwoArrExist .
func IsTwoArrExist(arr [][]int, target []int) bool {
	for _, v := range arr {
		if tool.IsTowSliceEqual(v, target) {
			fmt.Println(v, " ", target, " equal jump")
			return true
		}
	}
	return false
}

// IsArrExist .
func IsArrExist(arr []int, target int) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}
