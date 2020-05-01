package check

import "hwl/tool/logs"

// IsBoolEqual .
func IsBoolEqual(left, right bool) {
	if left == right {
		logs.Println("test ok")
	} else {
		logs.Println("test failed")
	}
}
