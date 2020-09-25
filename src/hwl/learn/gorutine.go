package learn

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// Goid 获取当前goroutine的id
// 在C/C++/Java等语言中，可以直接获取Thread Id。自从go1.4版本以后，Goroutine Id无法直接从Go Runtime获取了。
// 这是Golang的开发者故意为之，避免开发者滥用Goroutine Id实现Goroutine Local Storage(类似java的Thread Local Storage)，
// 因为Goroutine Local Storage很难进行垃圾回收。因此尽管Go1.4之前暴露出了相应的方法，现在已经把它隐藏了。
// 解决方法：将当前的堆栈信息写入到一个slice中，堆栈的第一行为 “goroutine #### […”，其中“####”就是当前的Goroutine Id
func Goid() int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic recover:panic info:%v", err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
