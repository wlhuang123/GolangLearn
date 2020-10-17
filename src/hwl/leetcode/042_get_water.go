package leetcode

import (
	"hwl/leetcode/check"
)

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。

示例:

输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6

*/

// Test042 .
func Test042() {
	Check042([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6)
	Check042([]int{}, 0)
	Check042([]int{1}, 0)
	Check042([]int{1, 1}, 0)
	Check042([]int{1, 6}, 0)
	Check042([]int{19, 6}, 0)
	Check042([]int{19, 0, 6}, 6)
}

// Check042 .
func Check042(height []int, expectResult int) {
	check.IsIntEqual(trap(height), expectResult)
}

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}

	var result int
	var leftIndex int
	// Check042([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6)
	rightIndex := len(height) - 1
	for leftIndex < rightIndex {
		if height[leftIndex] < height[rightIndex] {
			nextIndex, water := getNextIndexNoLowThanCur(height[leftIndex+1:rightIndex+1], height[leftIndex])
			result += water
			leftIndex += nextIndex + 1
			//fmt.Println("< index:", nextIndex, " water:", water, " leftIndex:", leftIndex)
			continue
		}

		nextIndex, water := getNextIndexNoLowThanCurRight(height[leftIndex:rightIndex], height[rightIndex])
		result += water
		rightIndex -= nextIndex
		//fmt.Println("index:", nextIndex, " water:", water, " rightIndex:", rightIndex)
	}

	//fmt.Println("result:", result)
	return result
}

// getNextIndexNoLowThanCur 获取下一个下标不低于当前下标高度的
// 如果能获取到下标，并返回这两个下标中间的储水量
func getNextIndexNoLowThanCur(height []int, curHight int) (int, int) {
	var water int
	for i, v := range height {
		if curHight > v {
			water += curHight - v
			continue
		}
		return i, water
	}
	return -1, water
}

// getNextIndexNoLowThanCurRight 获取下一个下标不低于当前下标高度的
// 如果能获取到下标，并返回这两个下标中间的储水量
func getNextIndexNoLowThanCurRight(height []int, curHight int) (int, int) {
	var water int
	for i := len(height) - 1; i >= 0; i-- {
		if curHight > height[i] {
			water += curHight - height[i]
			continue
		}
		return len(height) - i, water
	}
	return -1, water
}
