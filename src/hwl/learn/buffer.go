package learn

import (
	"bytes"
	"hwl/tool/logs"
)

// BufferTest .
func BufferTest() {
	readAndWriteBufferString()
}

func readAndWriteBufferString() {
	var b bytes.Buffer

	b.WriteString("first ")
	b.WriteString("second ")

	logs.Println(b.String())
}
