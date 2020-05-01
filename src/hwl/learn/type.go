package learn

import (
	"hwl/tool"
)

/*
   学习变量的类型
*/

// TypeTest .
func TypeTest() {
	stringType()
	byteType()
}

func stringType() {
	s := "abc"
	tool.GetType(s)    // string
	tool.GetType(s[0]) // uint8
}

func byteType() {
	b := []byte("abc")
	tool.GetType(b)    // 打印为空
	tool.GetType(b[0]) // uint8

	tool.GetType('{') // int32
}
