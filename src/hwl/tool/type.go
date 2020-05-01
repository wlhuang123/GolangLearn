package tool

import (
	"hwl/tool/logs"
	"reflect"
)

// GetType 获取变量的类型
func GetType(v interface{}) string {
	t := reflect.TypeOf(v).Name()
	logs.Println("type:", t)
	return t
}

// GetPointerAddr .
func GetPointerAddr(v interface{}) {
	logs.Printf("%p", v)
}
