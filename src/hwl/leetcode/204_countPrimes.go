package leetcode

import (
	"hwl/leetcode/check"
)

/*
统计所有小于非负整数 n 的质数的数量。

示例:

输入: 10
输出: 4
解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
*/

// Test204 .
func Test204() {
	Check204(10, 4)
	Check204(3, 1)
	Check204(4, 2)
	Check204(100000000, 0)
}

// Check204 .
func Check204(n, expectResult int) {
	check.IsIntEqual(countPrimes(n), expectResult)
}

func countPrimes(n int) int {
	// 方法1
	// var result int

	// if n < 3 {
	// 	return 0
	// }

	// for i := 3; i < n; i++ {
	// 	if isPrime(i) {
	// 		//fmt.Println(i)
	// 		result++
	// 	}
	// }

	// return result + 1

	// 方法2
	return getPrimeNum2(n)
}

func isPrime(n int) bool {
	if n%2 == 0 {
		return false
	}

	for i := 3; i*i < n+1; i = i + 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// getPrimeNum2 获取小于n的质数的方法2
// 把合数过滤掉，剩下的就是质数的数量
func getPrimeNum2(n int) int {
	slice := make([]int, n)
	if n < 3 {
		return 0
	}
	if n == 3 {
		return 1
	}

	for i := 2; i*i < n+1; i++ {
		for j := 2; ; j++ {
			temp := i * j
			if temp < n {
				slice[temp] = -1
				continue
			}
			break
		}
	}

	var result int
	for i := 2; i < len(slice); i++ {
		if slice[i] != -1 {
			result++
		}
	}
	return result
}
