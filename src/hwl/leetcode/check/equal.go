package check

import "hwl/tool/logs"

// IsBoolEqual .
func IsBoolEqual(left, right bool) {
	if left == right {
		logs.Println(OK)
	} else {
		logs.Println(Failed)
	}
}

// IsIntEqual .
func IsIntEqual(left, right int) {
	if left == right {
		logs.Println(OK)
	} else {
		logs.Println(Failed)
	}
}
