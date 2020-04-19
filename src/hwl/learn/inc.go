package learn

import "hwl/tool/logs"

/*
   自增++的介绍  --同理
*/

// IncTest /
func IncTest() {
	var i int
	i++ // ok 等价于 i=i+1
	logs.Println(i)
}
