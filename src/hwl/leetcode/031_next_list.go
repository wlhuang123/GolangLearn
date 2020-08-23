package leetcode

import (
	"hwl/leetcode/check"
	"hwl/tool"
)

/*
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/

// Test031 .
func Test031() {
	Check031([]int{1, 2, 3}, []int{1, 3, 2}) // 交换index 2
	Check031([]int{3, 2, 1}, []int{1, 2, 3}) //
	Check031([]int{1, 1, 5}, []int{1, 5, 1}) //
	Check031([]int{1, 3, 2}, []int{2, 1, 3}) // 1
	Check031([]int{1, 6, 5, 4}, []int{4, 1, 5, 6})
	Check031([]int{1, 5, 6, 4}, []int{1, 6, 4, 5})
	Check031([]int{9, 5, 7, 6}, []int{9, 6, 5, 7})
	Check031([]int{}, []int{})
	Check031([]int{6}, []int{6})
	Check031([]int{6, 7}, []int{7, 6})
	Check031([]int{7, 7}, []int{7, 7})
	Check031([]int{7, 6}, []int{6, 7})

	//1 8 9 6 5 4 3 2 7
}

// Check031 .
func Check031(nums, targetNums []int) {
	nextPermutation(nums)
	check.IsTowSliceTheSame(nums, targetNums)
}

// 尽量在最右边能满足交换条件的就交换
// 写一个函数，判断一个数组是否能交换，不能交换，就往前移
// 如何判断当前能否交换？
// 两个数字，判断后面比前面小，就交换
// 多个数字，此时后面数字肯定不能交换了，就看新加这个数字和之前哪个数字交换。找出能交换的数字，没有就继续往前移动，
// 有的话就找出能交换的最小那个数字，交换

// 问题：虽然能确定要交换的某个位置的数字，但是交换这个位置后，其他位置也要交换，所以应该是动态一直交换，而不是找到后交换一次
// 解决，确定要交换的位置后，此时前面的肯定都是降序排序的
// 交换的位置，右边开始最小那个
// 问题转化为一个数字加降序排列，如何变成升序
// 解决，先把该数字插入到正确位置，此时问题转化为如何将降序转为升序

// 231
// 132

func nextPermutation(nums []int) {
	for i := 0; i < len(nums); i++ {
		if CanExchange(nums[len(nums)-1-i:]) {
			return
		}
	}

	tool.OrderChange(nums)
}

// CanExchange 判断是否能交换,能交换就执行交换
func CanExchange(nums []int) bool {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[0] < nums[i] {
			nums[0], nums[i] = nums[i], nums[0]
			tool.OrderChange(nums[1:])
			return true
		}
	}
	return false
}
