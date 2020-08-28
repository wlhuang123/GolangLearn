package tool

// Max 返回较大值
func Max(left, right int) int {
	if left > right {
		return left
	}
	return right
}

// Least 返回较小值
func Least(left, right int) int {
	if left > right {
		return right
	}
	return left
}
