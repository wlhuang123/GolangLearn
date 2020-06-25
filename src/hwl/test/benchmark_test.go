package test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// fmt.Printf
func BenchmarkFmtSprintfMore(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s += fmt.Sprintf("%s%s", "hello", "world")
	}
}

// 加号 拼接
func BenchmarkAddMore(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s += "hello" + "world"
	}
}

// strings.Join
func BenchmarkStringsJoinMore(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s += strings.Join([]string{"hello", "world"}, "")
	}
}

// bytes.Buffer
func BenchmarkBufferMore(b *testing.B) {
	buffer := bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		buffer.WriteString("hello")
		buffer.WriteString("world")
	}
}
