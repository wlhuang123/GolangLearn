package learn

import (
	"hwl/tool/logs"
	"os"
)

/*
   命令行参数
*/

// ArgsTest .
// os.Args第一个元素是命令行本身 取决于启动程序的命令行
// 后面可以直接接参数 不限制个数 空格分开
func ArgsTest() {
	// eg: ./proc show pods
	for i := 0; i < len(os.Args); i++ {
		logs.Println("index:", i, " ", os.Args[i])
	}
}
