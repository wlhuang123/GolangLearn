package tool

// 获取系统的位数

// GetHostBits .
func GetHostBits() int {
	// 对uint(0)取反得到最大值
	// 32位右移63位得到0 ， 64位右移63位得到1
	return 32 << (^uint(0) >> 63)
}
