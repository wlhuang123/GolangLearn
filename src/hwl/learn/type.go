package learn

import (
	"fmt"
	"hwl/tool"
	"hwl/tool/logs"
	"strings"
	"unicode"
)

/*
   学习变量的类型
*/

// TypeTest .
func TypeTest() {
	stringType()
	byteType()
	runeType()
	rangeString()
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

// rune是int32的别名，几乎在所有方面等同于int32
// 它用来区分字符值和整数值
func runeType() {
	logs.Println("rune type:")
	var i rune
	tool.GetType(i)
}

// rangeString 遍历字符串 大写改成小写
func rangeString() {
	str := "HwlJrl"
	var newStr string

	for _, v := range str {
		if unicode.IsUpper(v) {
			newStr += strings.ToLower(string(v))
		} else {
			newStr += string(v)
		}
	}
	fmt.Println(newStr)
}
