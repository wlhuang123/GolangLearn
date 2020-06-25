package learn

import "hwl/tool/logs"

/*
   介绍copy函数
*/

// CopyTest .
func CopyTest() {
	copySlice()
}

func copySlice() {
	var s1 []string
	var s2 []string

	s1 = append(s1, "hello")

	copy(s2, s1)          // 其拷贝数据的长度为 len(dst)与len(src)之间的最小值，当s2为空时，不会拷贝
	logs.Println(len(s2)) // 0

	s2 = append(s2, "world1")
	s2 = append(s2, "world2")

	copy(s2, s1)
	logs.Println(len(s2)) // 2
	logs.Println(s2)      // hello world2
}
