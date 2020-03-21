package tool

import (
	"reflect"
)

// GetType 获取变量的类型
func GetType(v interface{}) string {
	return reflect.TypeOf(v).Name()
}
