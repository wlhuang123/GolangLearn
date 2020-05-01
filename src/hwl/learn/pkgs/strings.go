package pkgs

import (
	"hwl/tool/logs"
	"strings"
)

// StringsTest .
func StringsTest() {
	join()
	compare()
}

func join() {
	names := []string{"zhangsan", "lisi"}
	var s string
	for _, name := range names {
		// 代价比较高，做了一个字符串的拷贝，s需要垃圾回收。下面用Join的方式更高效
		s += name + " "
	}
	logs.Println(s)

	// 每个数组数组元素和空格连接，最后一个元素不和空格连接
	logs.Println(strings.Join(names, " "))
}

func compare() {
	// 定义里面注释说比较速度比===快
	strings.Compare("a", "b")
}
