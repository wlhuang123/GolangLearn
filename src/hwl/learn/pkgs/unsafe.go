package pkgs

import (
	"fmt"
	"unsafe"
)

// UnSafe .
type UnSafe struct {
}

// UnSafeTest .
func UnSafeTest() {
	var u UnSafe
	u.Sizeof()
}

// Sizeof .
func (u UnSafe) Sizeof() {
	var i int
	fmt.Println(unsafe.Sizeof(i)) // 8

	var s string
	fmt.Println(unsafe.Sizeof(s)) // 16 无论string有什么内容都是16

	var b bool
	fmt.Println(unsafe.Sizeof(b)) // 1

	var arr []int
	fmt.Println(unsafe.Sizeof(arr)) // 24 和数组里面的内容无关，固定24
}
