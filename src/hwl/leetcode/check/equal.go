package check

import (
	"hwl/tool"
	"hwl/tool/logs"
)

// IsBoolEqual .
func IsBoolEqual(left, right bool) {
	if left == right {
		logs.Println(OK)
	} else {
		logs.Println(Failed, " left:", left, " right:", right)
	}
}

// IsIntEqual .
func IsIntEqual(left, right int) {
	if left == right {
		logs.Println(OK)
	} else {
		logs.Println(Failed, " left:", left, " right:", right)
	}
}

// IsStringEqual .
func IsStringEqual(left, right string) {
	if left == right {
		logs.Println(OK)
	} else {
		logs.Println(Failed, " left:", left, " right:", right)
	}
}

// IsTowSliceTheSame 包括位置要相同
func IsTowSliceTheSame(left, right interface{}) {
	if tool.IsTowSliceTheSame(left, right) {
		logs.Println(OK)
		return
	}
	logs.Println(Failed, " left:", left, " right:", right)
}
