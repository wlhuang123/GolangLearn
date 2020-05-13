package learn

import "hwl/tool/logs"

// SliceTest .
func SliceTest() {
	funcSendSlice()
	nilsliceInit()
	emptySliceInit()
}

// emptySliceInit 初始化空切片
// 空切片用于函数返回告知获取的结果内容为空
func emptySliceInit() {
	s1 := make([]int, 0)
	if s1 != nil {
		logs.Println("s1 is not nil")
	}

	s2 := []int{}
	if s2 != nil {
		logs.Println("s2 is not nil")
	}
}

// nilsliceInit 初始化一个nil切片
// 当函数出现异常的时候返回
func nilsliceInit() {
	var s []string
	if s == nil {
		logs.Println("s is nil")
	}
}

// 函数间传递切片 函数内部修改切片会影响原始切片 因为拷贝的是切片的地址
func funcSendSlice() {
	var s []int
	s = append(s, 1)
	modifySlice(s)
	logs.Println("after modify s:", s)
}

func modifySlice(s []int) {
	s[0] = 2
	s = append(s, 502)
}
