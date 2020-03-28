package tool

import (
	"hwl/tool/logs"
	"reflect"
)

// GetType 获取变量的类型
func GetType(v interface{}) string {
	return reflect.TypeOf(v).Name()
}

// GetPointerAddr .
func GetPointerAddr(v interface{}) {
	logs.Printf("%p", v)
}
