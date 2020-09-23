package leetcode

import "hwl/leetcode/check"

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
*/

// Test014 .
func Test014() {
	Check014([]string{"ab", "avc"}, "a")
	Check014([]string{"flower", "flow", "flight"}, "fl")
	Check014([]string{"dog", "racecar", "car"}, "")
	Check014([]string{}, "")
	Check014([]string{""}, "")
}

// Check014 .
func Check014(strs []string, expectStr string) {
	check.IsStringEqual(longestCommonPrefix(strs), expectStr)
}

func longestCommonPrefix(strs []string) string {
	var result string

	if len(strs) == 0 {
		return ""
	}

	for index, c := range strs[0] {
		for i := 1; i < len(strs); i++ {
			if len(strs[i]) < index+1 {
				return result
			}

			temp := strs[i]
			if string(c) != string(temp[index]) {
				return result
			}
		}
		result += string(c)
	}

	return result
}
