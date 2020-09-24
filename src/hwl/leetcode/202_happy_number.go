package leetcode

import (
	"hwl/leetcode/check"
)

/*
编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。如果 可以变为  1，那么这个数就是快乐数。

如果 n 是快乐数就返回 True ；不是，则返回 False 。

示例：

输入：19
输出：true
解释：
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1

*/

// Test202 .
func Test202() {
	Check202(19, true)
	Check202(7, true)
	Check202(10, true)
}

// Check202 .
func Check202(n int, expectResult bool) {
	check.IsBoolEqual(isHappy(n), expectResult)
}

// 用一个全局map存储n，如果出现重复，就不是快乐数
func isHappy(n int) bool {
	slowN, fastN := n, n
	for {
		slowN = sum(slowN)
		if slowN == 1 {
			return true
		}

		fastN = sum(fastN)
		fastN = sum(fastN)

		if slowN == fastN {
			return false
		}
	}
}

func sum(n int) int {
	var newN int
	for n != 0 {
		endNum := n % 10
		newN += endNum * endNum
		n = n / 10
	}
	return newN
}
