package leetcode

import (
	"hwl/leetcode/check"
	"hwl/tool"
)

/*
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

*/

/*
分享：左右两端设置指针移动。比较左右 A B 两端高低，假设A比较低，如果A最终属于结果的一端，那么此时AB就是最大的结果，
A和其他节点的最大值=A*最大的底（即B）,此时记录下AB的最大值，A那边的指针往右移动，继续重复这个流程。
（B比A低同理）
*/

// Test011 .
func Test011() {
	Check011([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49)
}

// Check011 .
func Check011(height []int, result int) {
	check.IsIntEqual(maxArea(height), result)
}

func maxArea(height []int) int {
	var maxResult int

	begin := 0
	end := len(height) - 1
	maxResult = CalContain(height, begin, end)

	for {
		if begin >= end {
			break
		}

		var newResult int

		if height[begin] > height[end] {
			end--
		} else {
			begin++
		}
		newResult = CalContain(height, begin, end)

		if newResult > maxResult {
			maxResult = newResult
		}
	}

	return maxResult
}

// CalContain 计算容量
func CalContain(height []int, indexLeft, indexRight int) int {
	return (indexRight - indexLeft) * tool.Least(height[indexLeft], height[indexRight])
}
