package test

import (
	"hwl/tool/logs"
	"testing"
)

// BenchmarkAdd .
func BenchmarkAdd(b *testing.B) {
	logs.Println("hwl BenchMark:", b.N)
}
