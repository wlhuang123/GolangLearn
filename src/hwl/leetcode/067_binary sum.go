package leetcode

import (
	"hwl/leetcode/check"
)

/*
67. 二进制求和

给你两个二进制字符串，返回它们的和（用二进制表示）。
输入为 非空 字符串且只包含数字 1 和 0。

示例 1:
输入: a = "11", b = "1"
输出: "100"

示例 2:
输入: a = "1010", b = "1011"
输出: "10101"
*/

// Test067 .
func Test067() {
	Check067("11", "1", "100")
	Check067("1010", "1011", "10101")
}

// Check067 .
func Check067(a, b string, expect string) {
	check.IsStringEqual(addBinary(a, b), expect)
}

func addBinary(a string, b string) string {
	var res string
	indexa := len(a) - 1
	indexb := len(b) - 1

	var isAdd uint8
	var tempa, tempb uint8
	for {
		tempa, tempb = 0, 0
		if indexa < 0 && indexb < 0 {
			break
		}
		if indexa >= 0 {
			tempa = uint8(a[indexa] - '0')
			indexa--
		}
		if indexb >= 0 {
			tempb = uint8(b[indexb] - '0')
			indexb--
		}
		var tempRes string
		tempRes, isAdd = getResult067(tempa, tempb, isAdd)
		res = tempRes + res
	}
	if isAdd == 1 {
		res = "1" + res
	}
	return res
}

func getResult067(a, b, isAdd uint8) (string, uint8) {
	var count uint8

	count = a + b + isAdd
	if count == 0 {
		return "0", 0
	}
	if count == 1 {
		return "1", 0
	}
	if count == 2 {
		return "0", 1
	}
	return "1", 1
}
