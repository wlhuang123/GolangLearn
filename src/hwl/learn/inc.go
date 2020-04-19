package learn

import "hwl/tool/logs"

/*
   自增++的介绍  --同理
*/

// IncTest /
func IncTest() {
	var i int
	i++ // ok 等价于 i=i+1
	// num:=i++    not ok, i++是一个语句，不像C语言是一个表达式，不可以赋值
	// ++i not ok，只支持后置
	logs.Println(i)
}
