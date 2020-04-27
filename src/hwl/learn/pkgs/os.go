package pkgs

import (
	"hwl/tool/logs"
	"os"
)

/*
   os包介绍：主要是为了以平台无关的方式和操作系统打交道
*/

// OSTest .
func OSTest() {
	fileTest()
}

// fileTest 文件句柄指针相关
func fileTest() {
	// 关闭一个打开文件失败的句柄， 关闭不会导致程序崩溃，Close内部做了nil指针判断
	f, _ := os.Open("qpz")
	if err := f.Close(); err != nil {
		logs.Println(err)
	}

	// 重复关闭一个打开的文件句柄，不会导致程序崩溃，会报错提示ile already closed
	f, _ = os.Open("/proc/1/stat")
	if err := f.Close(); err != nil {
		logs.Println(err)
	}
	if err := f.Close(); err != nil {
		logs.Println(err)
	}
}
