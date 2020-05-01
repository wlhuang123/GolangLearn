package leetcode

import (
	"hwl/leetcode/check"
)

// Test007 .
func Test007() {
	Check007(reverse(-1), -1)
	Check007(reverse(1), 1)
	Check007(reverse(0), 0)
	Check007(reverse(10), 1)
	Check007(reverse(123), 321)
	Check007(reverse(-123), -321)
	Check007(reverse(1534236469), 0)
}

// Check007 .
func Check007(left, right int) {
	check.IsIntEqual(left, right)
}

func reverse(x int) int {
	y := 0
	for x != 0 {
		y = y*10 + x%10
		if y < (-1<<31) || y > (1<<31)-1 {
			return 0
		}
		x /= 10
	}
	return y
}
