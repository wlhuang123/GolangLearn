package link

import _ "unsafe"

//go:linkname helloWorld hwl/demo/hello.Hello
func helloWorld() {
	println("hello world!")
}
